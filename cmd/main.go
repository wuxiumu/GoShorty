// cmd/main.go - 服务入口
package main

import (
	"fmt"
	"log"
	"net/http"

	"goshor/internal/config"
	"goshor/internal/core"
	"goshor/internal/handler"
	"goshor/internal/pprof"
	"goshor/internal/util"
)

func main() {
	cfg := config.Load()
	core.InitLimiter() // ✅ 初始化令牌桶限流器
	util.StartLogger() // ✅ 启动日志记录器
	pprof.Start(cfg.MetricsAddr)

	// 公开接口使用独立 ServeMux。
	// net/http/pprof 在 init 中注册到 DefaultServeMux，指标服务复用该 Mux，避免调试端点出现在对外端口。
	mux := http.NewServeMux()
	mux.HandleFunc("/api/shorten", handler.ShortenHandler)
	mux.HandleFunc("/api/stats/", handler.StatsHandler)
	mux.HandleFunc("/", handler.RedirectHandler)

	fmt.Printf("🚀 GoShorty HTTP 监听：%s\n", cfg.HTTPAddr)
	fmt.Printf("📈 pprof / metrics 监听：%s\n", cfg.MetricsAddr)
	log.Fatal(http.ListenAndServe(cfg.HTTPAddr, mux))
}
