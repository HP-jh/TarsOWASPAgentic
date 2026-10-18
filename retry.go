package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

// RetryWithBackoff 带指数退避的重试
func RetryWithBackoff(ctx context.Context, maxRetries int, fn func() error) error {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		if err := fn(); err == nil {
			return nil
		} else {
			lastErr = err
		}
		if i < maxRetries-1 {
			delay := time.Duration(math.Pow(2, float64(i))) * time.Second
			if delay > 10*time.Second {
				delay = 10 * time.Second
			}
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
	return fmt.Errorf("重试 %d 次后仍失败: %w", maxRetries, lastErr)
}
