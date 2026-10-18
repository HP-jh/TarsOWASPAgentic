package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type AuditEntry struct {
	Time    time.Time `json:"time"`
	Tool    string    `json:"tool"`
	Target  string    `json:"target,omitempty"`
	Status  string    `json:"status"`
	Error   string    `json:"error,omitempty"`
}

type AuditLog struct {
	mu      sync.Mutex
	entries []AuditEntry
	file    string
}

func NewAuditLog(path string) *AuditLog {
	return &AuditLog{file: path}
}

func (a *AuditLog) Record(tool, target, status, errMsg string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	entry := AuditEntry{Time: time.Now().UTC(), Tool: tool, Target: target, Status: status, Error: errMsg}
	a.entries = append(a.entries, entry)
	if a.file != "" {
		data, _ := json.Marshal(entry)
		f, _ := os.OpenFile(a.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if f != nil {
			fmt.Fprintln(f, string(data))
			f.Close()
		}
	}
}

func (a *AuditLog) Entries() []AuditEntry {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]AuditEntry, len(a.entries))
	copy(out, a.entries)
	return out
}
