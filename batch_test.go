package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBatchScan(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/scan", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"status": "ok"})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()
	client := NewClient(srv.URL, "")
	results := BatchScan(context.Background(), client, []string{"a.com", "b.com"}, 2)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if results[0].Error != "" {
		t.Fatalf("unexpected error: %s", results[0].Error)
	}
}

func TestBatchHandler_MethodNotAllowed(t *testing.T) {
	client := NewClient("", "")
	req := httptest.NewRequest(http.MethodGet, "/batch", nil)
	w := httptest.NewRecorder()
	withClient(client, batchHandler)(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}

func TestBatchHandler_EmptyTargets(t *testing.T) {
	client := NewClient("", "")
	body, _ := json.Marshal(map[string]interface{}{"targets": []string{}})
	req := httptest.NewRequest(http.MethodPost, "/batch", bytes.NewReader(body))
	w := httptest.NewRecorder()
	withClient(client, batchHandler)(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestBatchHandler_TooManyTargets(t *testing.T) {
	client := NewClient("", "")
	targets := make([]string, 51)
	body, _ := json.Marshal(map[string]interface{}{"targets": targets})
	req := httptest.NewRequest(http.MethodPost, "/batch", bytes.NewReader(body))
	w := httptest.NewRecorder()
	withClient(client, batchHandler)(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
