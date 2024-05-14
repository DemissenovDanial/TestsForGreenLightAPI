package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotFoundResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	app.notFoundResponse(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d; got %d", http.StatusNotFound, rr.Code)
	}

	expectedBody := `{"error":"the requested resource could not be found"}`
	actualBody := strings.TrimSpace(rr.Body.String())
	actualBody = strings.Replace(actualBody, "\n", "", -1)
	actualBody = strings.Replace(actualBody, "\t", "", -1)
	actualBody = strings.Replace(actualBody, ": ", ":", -1)
	if actualBody != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, actualBody)
	}
}
