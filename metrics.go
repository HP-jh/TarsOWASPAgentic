package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Metrics struct {
	RequestsTotal   uint64
	RequestsSuccess uint64
	RequestsFail    uint64
	ScanTotal       uint64
	ScanSuccess     uint64
	ScanFail        uint64
	StartTime       time.Time
}

var globalMetrics = &Metrics{StartTime: time.Now()}
var muMetrics sync.RWMutex

func IncRequest(success bool) {
	atomic.AddUint64(&globalMetrics.RequestsTotal, 1)
	if success {
		atomic.AddUint64(&globalMetrics.RequestsSuccess, 1)
	} else {
		atomic.AddUint64(&globalMetrics.RequestsFail, 1)
	}
}

func IncScan(success bool) {
	atomic.AddUint64(&globalMetrics.ScanTotal, 1)
	if success {
		atomic.AddUint64(&globalMetrics.ScanSuccess, 1)
	} else {
		atomic.AddUint64(&globalMetrics.ScanFail, 1)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request, client *Client) {
	muMetrics.RLock()
	defer muMetrics.RUnlock()
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"uptime_seconds": time.Since(globalMetrics.StartTime).Seconds(),
		"requests": map[string]uint64{
			"total":   atomic.LoadUint64(&globalMetrics.RequestsTotal),
			"success": atomic.LoadUint64(&globalMetrics.RequestsSuccess),
			"fail":    atomic.LoadUint64(&globalMetrics.RequestsFail),
		},
		"scans": map[string]uint64{
			"total":   atomic.LoadUint64(&globalMetrics.ScanTotal),
			"success": atomic.LoadUint64(&globalMetrics.ScanSuccess),
			"fail":    atomic.LoadUint64(&globalMetrics.ScanFail),
		},
	})
}
