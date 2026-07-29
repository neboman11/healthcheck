package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewServerServesHealthcheck(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set("Origin", "https://example.com")
	recorder := httptest.NewRecorder()

	NewServer().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if body := recorder.Body.String(); body != "Health check successful." {
		t.Fatalf("body = %q, want %q", body, "Health check successful.")
	}
	if origin := recorder.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Fatalf("Access-Control-Allow-Origin = %q, want %q", origin, "*")
	}
}

func TestNewServerHandlesHealthcheckPreflight(t *testing.T) {
	request := httptest.NewRequest(http.MethodOptions, "/", nil)
	request.Header.Set("Origin", "https://example.com")
	request.Header.Set("Access-Control-Request-Method", http.MethodGet)
	recorder := httptest.NewRecorder()

	NewServer().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusNoContent)
	}
	if origin := recorder.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Fatalf("Access-Control-Allow-Origin = %q, want %q", origin, "*")
	}
	if methods := recorder.Header().Get("Access-Control-Allow-Methods"); !strings.Contains(methods, http.MethodGet) {
		t.Fatalf("Access-Control-Allow-Methods = %q, want it to include %q", methods, http.MethodGet)
	}
}
