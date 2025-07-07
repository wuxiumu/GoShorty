// cmd/main.go - 服务入口
package main

import (
	"fmt"
	"goshor/internal/core"
	"goshor/internal/handler"
	"goshor/internal/util"
	"log"
	"net/http"

	_ "goshor/internal/pprof"
	_ "net/http/pprof" // ✅ 集成 pprof 分析器
)

func main() {
	core.InitLimiter() // ✅ 初始化令牌桶限流器
	util.StartLogger() // ✅ 启动日志记录器

	http.HandleFunc("/api/shorten", handler.ShortenHandler)
	http.HandleFunc("/api/stats/", handler.StatsHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	fmt.Println("🚀 GoShorty 服务启动：http://localhost:8080")
	fmt.Println("📈 pprof 可访问：http://localhost:8080/debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
