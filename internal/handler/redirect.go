// internal/handler/redirect.go
package handler

import (
	"fmt"
	"net/http"
	"strings"

	"goshor/internal/core"
	"goshor/internal/pprof"
	"goshor/internal/util"
)

// RedirectHandler 处理跳转逻辑，使用令牌桶限流
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	if !core.Allow() {
		pprof.LimitedCounter.Inc()
		util.AsyncLog("[LIMIT] 被限流")
		http.Error(w, "Too many requests", http.StatusTooManyRequests)
		return
	}

	key := strings.TrimPrefix(r.URL.Path, "/")
	core.StoreMu.Lock()
	link, ok := core.Store[key]
	if ok {
		pprof.VisitCounter.Inc()
		link.Visits++
		util.AsyncLog(fmt.Sprintf("[VISIT] %s => %s", key, link.Original))
	}
	core.StoreMu.Unlock()

	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, link.Original, http.StatusFound)
}
