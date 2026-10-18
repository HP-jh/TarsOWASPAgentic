package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"net/http"
	"time"
)

// BatchScanResult 批量扫描单条结果
type BatchScanResult struct {
	Target string                 `json:"target"`
	Result map[string]interface{} `json:"result,omitempty"`
	Error  string                 `json:"error,omitempty"`
}

// BatchScan 并发批量扫描多个目标
func BatchScan(ctx context.Context, client *Client, targets []string, maxConcurrency int) []BatchScanResult {
	if maxConcurrency <= 0 {
		maxConcurrency = 3
	}
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	results := make([]BatchScanResult, len(targets))
	var mu sync.Mutex

	for i, target := range targets {
		wg.Add(1)
		go func(idx int, t string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			res, err := client.req(ctx, "POST", "/api/v1/scan", map[string]interface{}{"target": t})
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				results[idx] = BatchScanResult{Target: t, Error: FriendlyError(err)}
			} else {
				results[idx] = BatchScanResult{Target: t, Result: res}
			}
		}(i, target)
	}
	wg.Wait()
	return results
}

func batchHandler(w http.ResponseWriter, r *http.Request, client *Client) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "只接受 POST 请求"})
		return
	}
	var req struct {
		Targets        []string `json:"targets"`
		MaxConcurrency int      `json:"max_concurrency"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": fmt.Sprintf("解析失败：%v", err)})
		return
	}
	if len(req.Targets) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "targets 不能为空"})
		return
	}
	if len(req.Targets) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": "单次批量最多 50 个目标"})
		return
	}
	for _, t := range req.Targets {
		if err := ValidateTarget(t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"error": fmt.Sprintf("目标校验失败 (%s): %v", t, err)})
			return
		}
	}
	ctx, cancel := context.WithTimeout(r.Context(), 300*time.Second)
	defer cancel()
	results := BatchScan(ctx, client, req.Targets, req.MaxConcurrency)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"results": results})
}
