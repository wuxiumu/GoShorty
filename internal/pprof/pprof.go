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
	http.Handle("/metrics", promhttp.Handler()) // 注册 prometheus 的 handler
	go func() {
		log.Println("[pprof] running at http://localhost:6060/debug/pprof/")
		log.Println("[metrics] running at http://localhost:6060/metrics") // 启动 prometheus 的 http 服务
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
