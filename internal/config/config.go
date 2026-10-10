// Package config 读取进程环境变量，并在变量未设置时补上工作目录里的 .env。
// 示例见仓库根目录 .env.example。这里没有密钥或第三方令牌。
package config

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

// Config 是进程启动时确定的监听地址与对外基址。
type Config struct {
	HTTPAddr      string
	MetricsAddr   string
	PublicBaseURL string
}

var (
	mu      sync.RWMutex
	current = Config{
		HTTPAddr:      ":8080",
		MetricsAddr:   "127.0.0.1:6060",
		PublicBaseURL: "http://localhost:8080",
	}
)

// Load 解析当前环境。已存在的环境变量优先于 .env 文件。
func Load() Config {
	file := readDotEnv(".env")
	cfg := Config{
		HTTPAddr:      firstNonEmpty(os.Getenv("HTTP_ADDR"), file["HTTP_ADDR"], ":8080"),
		MetricsAddr:   firstNonEmpty(os.Getenv("METRICS_ADDR"), file["METRICS_ADDR"], "127.0.0.1:6060"),
		PublicBaseURL: strings.TrimRight(firstNonEmpty(os.Getenv("PUBLIC_BASE_URL"), file["PUBLIC_BASE_URL"], "http://localhost:8080"), "/"),
	}
	mu.Lock()
	current = cfg
	mu.Unlock()
	return cfg
}

// Current 返回最近一次 Load 的结果；未调用 Load 时是代码内默认值。
func Current() Config {
	mu.RLock()
	defer mu.RUnlock()
	return current
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func readDotEnv(path string) map[string]string {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	out := make(map[string]string)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(strings.TrimPrefix(line, "export "))
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if key == "" || strings.ContainsAny(key, " \t#") {
			continue
		}
		val = strings.TrimSpace(val)
		if len(val) >= 2 {
			if (val[0] == '"' && val[len(val)-1] == '"') || (val[0] == '\'' && val[len(val)-1] == '\'') {
				val = val[1 : len(val)-1]
			}
		}
		out[key] = val
	}
	return out
}
