package main

import (
	"testing"
)

func TestValidateTarget_Empty(t *testing.T) {
	if err := ValidateTarget(""); err == nil {
		t.Fatal("expected error for empty target")
	}
}

func TestValidateTarget_ValidURL(t *testing.T) {
	if err := ValidateTarget("https://example.com"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateTarget_InvalidScheme(t *testing.T) {
	if err := ValidateTarget("ftp://example.com"); err == nil {
		t.Fatal("expected error for ftp")
	}
}

func TestValidateTarget_LocalhostBlocked(t *testing.T) {
	if err := ValidateTarget("http://localhost:8080"); err == nil {
		t.Fatal("expected error for localhost")
	}
}

func TestValidateTarget_PrivateIPBlocked(t *testing.T) {
	if err := ValidateTarget("http://192.168.1.1"); err == nil {
		t.Fatal("expected error for private IP")
	}
}
