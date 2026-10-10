// internal/pprof/pprof.go
package pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	VisitCounter   prometheus.Counter
	LimitedCounter prometheus.Counter
)

// 自动挂载 pprof 的 init 函数
func init() {
	VisitCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "shorty_total_visits",
		Help: "Total number of successful short link redirects",
	})
	LimitedCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "shorty_limited_requests",
		Help: "Total number of requests that were rate-limited",
	})
	prometheus.MustRegister(VisitCounter)
	prometheus.MustRegister(LimitedCounter)
	http.Handle("/metrics", promhttp.Handler()) // 注册到 DefaultServeMux，仅由 Start 监听
}

// Start 在 addr 上提供 pprof 与 /metrics。调用方应使用独立 Mux 提供对外 HTTP。
func Start(addr string) {
	go func() {
		log.Printf("[pprof] http://%s/debug/pprof/", addr)
		log.Printf("[metrics] http://%s/metrics", addr)
		log.Println(http.ListenAndServe(addr, nil))
	}()
}
