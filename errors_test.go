package main

import (
	"errors"
	"testing"
)

func TestFriendlyError_Timeout(t *testing.T) {
	msg := FriendlyError(errors.New("context deadline exceeded"))
	if msg != "连接超时，请检查网络或稍后重试" {
		t.Fatalf("unexpected friendly msg: %s", msg)
	}
}

func TestFriendlyError_Refused(t *testing.T) {
	msg := FriendlyError(errors.New("connection refused"))
	if msg != "目标服务未启动或无法访问" {
		t.Fatalf("unexpected friendly msg: %s", msg)
	}
}

func TestFriendlyError_Nil(t *testing.T) {
	msg := FriendlyError(nil)
	if msg != "" {
		t.Fatalf("expected empty, got %s", msg)
	}
}

func TestFriendlyError_Generic(t *testing.T) {
	msg := FriendlyError(errors.New("something weird"))
	if msg == "" {
		t.Fatal("expected non-empty generic msg")
	}
}
