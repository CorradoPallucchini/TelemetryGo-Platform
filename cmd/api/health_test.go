package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	rr := httptest.NewRecorder()

	healthCheckHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Status code errato: ottenuto %v, atteso %v", rr.Code, http.StatusOK)
	}

	expectedBody := `{"status": "ok"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Body errato: ottenuto %v, atteso %v", rr.Body.String(), expectedBody)
	}
}
