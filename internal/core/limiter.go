package core

import (
	"sync"
	"time"
)

// TokenBucket implements a basic token bucket rate limiter
type TokenBucket struct {
	capacity int           // max tokens in bucket
	tokens   int           // current token count
	interval time.Duration // interval to add tokens
	last     time.Time     // last refill time
	mu       sync.Mutex
}

var bucket *TokenBucket

// InitLimiter initializes the token bucket limiter
func InitLimiter() {
	bucket = &TokenBucket{
		capacity: 10,
		tokens:   10,
		interval: 100 * time.Millisecond,
		last:     time.Now(),
	}
}

// Allow returns true if a request is allowed to proceed
func Allow() bool {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(bucket.last)
	regen := int(elapsed / bucket.interval)

	if regen > 0 {
		bucket.tokens = min(bucket.capacity, bucket.tokens+regen)
		bucket.last = now
	}

	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}
	return false
}

// 如需加入更复杂的策略（如按 IP 分组限流），也可以继续扩展此模块。
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
