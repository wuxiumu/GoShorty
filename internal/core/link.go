// internal/core/link.go
package core

import (
	"fmt"
	"sync"
	"time"
)

// Link 短链接实体结构（内存对齐优化）
type Link struct {
	CreatedAt time.Time
	Original  string
	Visits    int64
}

// 全局存储与并发控制
var (
	Store   = make(map[string]*Link)
	StoreMu sync.RWMutex
	LimitCh = make(chan struct{}, 10) // 限流通道
)

// GenerateShortKey 生成短 key
func GenerateShortKey(url string) string {
	return fmt.Sprintf("%x", time.Now().UnixNano())[:6]
}
