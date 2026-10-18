package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	tokens   int
	max      int
	interval time.Duration
	last     time.Time
}

func NewRateLimiter(max int, interval time.Duration) *RateLimiter {
	return &RateLimiter{tokens: max, max: max, interval: interval, last: time.Now()}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(rl.last)
	refill := int(elapsed / rl.interval)
	if refill > 0 {
		rl.tokens += refill
		if rl.tokens > rl.max {
			rl.tokens = rl.max
		}
		rl.last = now
	}
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func RateLimitMiddleware(rl *RateLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "请求太频繁，请稍后再试"})
			return
		}
		next(w, r)
	}
}
