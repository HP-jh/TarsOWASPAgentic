package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIncRequest(t *testing.T) {
	globalMetrics.RequestsTotal = 0
	globalMetrics.RequestsSuccess = 0
	globalMetrics.RequestsFail = 0
	IncRequest(true)
	IncRequest(false)
	if globalMetrics.RequestsTotal != 2 {
		t.Fatalf("expected 2 total, got %d", globalMetrics.RequestsTotal)
	}
	if globalMetrics.RequestsSuccess != 1 {
		t.Fatalf("expected 1 success, got %d", globalMetrics.RequestsSuccess)
	}
}

func TestMetricsHandler(t *testing.T) {
	client := NewClient("", "")
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()
	metricsHandler(w, req, client)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if _, ok := resp["uptime_seconds"]; !ok {
		t.Fatal("expected uptime_seconds")
	}
}
