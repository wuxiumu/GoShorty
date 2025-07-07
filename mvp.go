// GoShorty 短链接服务入口 + 高级特性演示 + 性能分析支持
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // ✅ 启用 pprof 分析器
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

// Link 表示一个短链接的数据结构（内存对齐优化）
type Link struct {
	CreatedAt time.Time
	Original  string
	Visits    int64
}

var (
	store   = make(map[string]*Link)
	storeMu sync.RWMutex
	logChan = make(chan string, 100)
	limitCh = make(chan struct{}, 10) // ✅ 限流器：最多允许10个并发跳转
)

func generateShortKey(url string) string {
	return fmt.Sprintf("%x", time.Now().UnixNano())[:6]
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	key := generateShortKey(url)

	storeMu.Lock()
	store[key] = &Link{Original: url, CreatedAt: time.Now()}
	storeMu.Unlock()

	select {
	case logChan <- fmt.Sprintf("[NEW] %s => %s", key, url):
	default:
	}

	resp := map[string]string{"short": fmt.Sprintf("http://localhost:8080/%s", key)}
	_ = json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case limitCh <- struct{}{}:
		defer func() { <-limitCh }()
	default:
		http.Error(w, "Too many concurrent requests", http.StatusTooManyRequests)
		return
	}

	key := strings.TrimPrefix(r.URL.Path, "/")
	storeMu.Lock()
	link, ok := store[key]
	if ok {
		link.Visits++
		select {
		case logChan <- fmt.Sprintf("[VISIT] %s => %s", key, link.Original):
		default:
		}
	}
	storeMu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/api/stats/")
	storeMu.RLock()
	link, ok := store[key]
	storeMu.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	_ = json.NewEncoder(w).Encode(link)
}

func exportCSV(data map[string]*Link, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	t := reflect.TypeOf(Link{})
	headers := make([]string, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		headers[i] = t.Field(i).Name
	}
	_ = w.Write(headers)

	for _, v := range data {
		r := reflect.ValueOf(*v)
		row := make([]string, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			row[i] = fmt.Sprintf("%v", r.Field(i).Interface())
		}
		_ = w.Write(row)
	}
	return nil
}

func MapReduce[T any, R any](list []T, mapper func(T) R, reducer func(R, R) R, zero R) R {
	var result = zero
	for _, item := range list {
		result = reducer(result, mapper(item))
	}
	return result
}

func startLogger() {
	go func() {
		for logMsg := range logChan {
			log.Println(logMsg)
		}
	}()
}

func main() {
	startLogger()

	http.HandleFunc("/api/shorten", shortenHandler)
	http.HandleFunc("/api/stats/", statsHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("🚀 GoShorty 服务启动：http://localhost:8080")
	fmt.Println("📈 pprof 可访问：http://localhost:8080/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
