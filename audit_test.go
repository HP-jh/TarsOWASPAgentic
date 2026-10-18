package main

import (
	"os"
	"testing"
)

func TestAuditLog_RecordAndEntries(t *testing.T) {
	al := NewAuditLog("")
	al.Record("owasp_scan", "example.com", "success", "")
	al.Record("owasp_status", "", "fail", "timeout")
	entries := al.Entries()
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if entries[0].Tool != "owasp_scan" {
		t.Fatalf("expected owasp_scan, got %s", entries[0].Tool)
	}
	if entries[1].Status != "fail" {
		t.Fatalf("expected fail, got %s", entries[1].Status)
	}
}

func TestAuditLog_FileWrite(t *testing.T) {
	f, _ := os.CreateTemp("", "audit*.log")
	defer os.Remove(f.Name())
	f.Close()
	al := NewAuditLog(f.Name())
	al.Record("owasp_scan", "test.com", "success", "")
	data, err := os.ReadFile(f.Name())
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("expected audit file to have content")
	}
}
