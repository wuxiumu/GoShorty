// internal/handler/stats.go
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"goshor/internal/core"
)

// StatsHandler 返回统计信息
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/api/stats/")
	core.StoreMu.RLock()
	link, ok := core.Store[key]
	core.StoreMu.RUnlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	_ = json.NewEncoder(w).Encode(link)
}
