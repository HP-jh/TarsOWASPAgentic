package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimiter_Allow(t *testing.T) {
	rl := NewRateLimiter(2, time.Second)
	if !rl.Allow() {
		t.Fatal("expected first allow")
	}
	if !rl.Allow() {
		t.Fatal("expected second allow")
	}
	if rl.Allow() {
		t.Fatal("expected third deny")
	}
}

func TestRateLimiter_Refill(t *testing.T) {
	rl := NewRateLimiter(1, 50*time.Millisecond)
	if !rl.Allow() {
		t.Fatal("expected first allow")
	}
	if rl.Allow() {
		t.Fatal("expected deny before refill")
	}
	time.Sleep(60 * time.Millisecond)
	if !rl.Allow() {
		t.Fatal("expected allow after refill")
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	rl := NewRateLimiter(1, time.Hour)
	handler := RateLimitMiddleware(rl, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w1 := httptest.NewRecorder()
	handler(w1, req)
	if w1.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w1.Code)
	}
	w2 := httptest.NewRecorder()
	handler(w2, req)
	if w2.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429, got %d", w2.Code)
	}
}
