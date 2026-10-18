package main

import (
	"os"
	"testing"
)

func TestLoadConfigFromFile_Valid(t *testing.T) {
	f, _ := os.CreateTemp("", "config*.json")
	defer os.Remove(f.Name())
	f.WriteString(`{"endpoint":"https://test.com","api_key":"k","mode":"http","timeout_seconds":10}`)
	f.Close()
	cfg, err := LoadConfigFromFile(f.Name())
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if cfg.Endpoint != "https://test.com" {
		t.Fatalf("expected https://test.com, got %s", cfg.Endpoint)
	}
}

func TestLoadConfigFromFile_Defaults(t *testing.T) {
	f, _ := os.CreateTemp("", "config*.json")
	defer os.Remove(f.Name())
	f.WriteString(`{}`)
	f.Close()
	cfg, err := LoadConfigFromFile(f.Name())
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if cfg.Endpoint != "https://api.example.com" {
		t.Fatalf("expected default endpoint, got %s", cfg.Endpoint)
	}
	if cfg.Timeout != 30 {
		t.Fatalf("expected default timeout 30, got %d", cfg.Timeout)
	}
}

func TestLoadConfigFromFile_Missing(t *testing.T) {
	_, err := LoadConfigFromFile("/nonexistent/config.json")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
