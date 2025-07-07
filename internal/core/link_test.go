package core

import (
	"testing"
	"time"
)

func TestGenerateShortKey(t *testing.T) {
	key := GenerateShortKey("https://golang.org")
	if len(key) != 6 {
		t.Errorf("expected key length 6, got %d", len(key))
	}
}

func TestTokenBucketLimiter(t *testing.T) {
	InitLimiter()

	allowed := 0
	for i := 0; i < 20; i++ {
		if Allow() {
			allowed++
		}
	}
	if allowed > 10 {
		t.Errorf("expected at most 10 allowed, got %d", allowed)
	}

	time.Sleep(500 * time.Millisecond)

	// after sleep, should be able to allow a few more
	refill := 0
	for i := 0; i < 5; i++ {
		if Allow() {
			refill++
		}
	}
	if refill == 0 {
		t.Error("expected tokens to be refilled after sleep")
	}
}
