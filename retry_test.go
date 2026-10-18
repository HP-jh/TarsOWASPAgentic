package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryWithBackoff_SuccessFirst(t *testing.T) {
	called := 0
	err := RetryWithBackoff(context.Background(), 3, func() error {
		called++
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if called != 1 {
		t.Fatalf("expected 1 call, got %d", called)
	}
}

func TestRetryWithBackoff_SuccessAfterTwo(t *testing.T) {
	called := 0
	err := RetryWithBackoff(context.Background(), 3, func() error {
		called++
		if called < 2 {
			return errors.New("temp fail")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if called != 2 {
		t.Fatalf("expected 2 calls, got %d", called)
	}
}

func TestRetryWithBackoff_Exhausted(t *testing.T) {
	called := 0
	err := RetryWithBackoff(context.Background(), 2, func() error {
		called++
		return errors.New("always fail")
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if called != 2 {
		t.Fatalf("expected 2 calls, got %d", called)
	}
}

func TestRetryWithBackoff_ContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := RetryWithBackoff(ctx, 3, func() error {
		return errors.New("fail")
	})
	if err != context.Canceled {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}

func TestRetryWithBackoff_MaxDelay(t *testing.T) {
	start := time.Now()
	_ = RetryWithBackoff(context.Background(), 5, func() error {
		return errors.New("fail")
	})
	// 2^0 + 2^1 + 2^2 + 2^3 capped at 10s each, total < 20s
	if time.Since(start) > 20*time.Second {
		t.Fatal("backoff took too long")
	}
}
