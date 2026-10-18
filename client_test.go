package main

import (
	"testing"
	"time"
)

func TestNewOptimizedClient(t *testing.T) {
	c := NewOptimizedClient(30 * time.Second)
	if c == nil {
		t.Fatal("expected non-nil client")
	}
	if c.Timeout != 30*time.Second {
		t.Fatalf("expected 30s timeout, got %v", c.Timeout)
	}
	if c.Transport == nil {
		t.Fatal("expected transport")
	}
}
