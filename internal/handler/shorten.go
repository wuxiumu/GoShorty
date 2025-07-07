// internal/handler/shorten.go
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"goshor/internal/core"
	"goshor/internal/util"
)

// ShortenHandler 处理短链接生成请求
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	key := core.GenerateShortKey(url)

	core.StoreMu.Lock()
	core.Store[key] = &core.Link{Original: url, CreatedAt: time.Now()}
	core.StoreMu.Unlock()

	util.AsyncLog(fmt.Sprintf("[NEW] %s => %s", key, url))

	resp := map[string]string{"short": fmt.Sprintf("http://localhost:8080/%s", key)}
	_ = json.NewEncoder(w).Encode(resp)
}
