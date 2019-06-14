package main

import (
	"net/http"
	"testing"
)

func TestDetermineIP(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = "1.1.1.1:12345"

	if ip := DetermineIP(req); ip != "1.1.1.1" {
		t.Errorf("RemoteAddr IP is wrong: got %s, want %s", ip, "1.1.1.1")
	}

	req.Header.Set("X-Real-IP", "2.2.2.2")

	if ip := DetermineIP(req); ip != "2.2.2.2" {
		t.Errorf("X-Real-IP IP is wrong: got %s, want %s", ip, "2.2.2.2")
	}

	req.Header.Set("X-Forwarded-For", "3.3.3.3")

	if ip := DetermineIP(req); ip != "3.3.3.3" {
		t.Errorf("X-Forwarded-For IP is wrong: got %s, want %s", ip, "3.3.3.3")
	}
}
