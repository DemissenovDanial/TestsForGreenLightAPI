package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBadRequestResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	err := fmt.Errorf("test error")
	app.badRequestResponse(rr, req, err)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d; got %d", http.StatusBadRequest, rr.Code)
	}

	expectedBody := `{"error":"test error"}`
	actualBody := rr.Body.String()
	actualBody = strings.Replace(actualBody, "\n", "", -1)
	actualBody = strings.Replace(actualBody, "\t", "", -1)
	actualBody = strings.Replace(actualBody, ": ", ":", -1)
	if actualBody != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, actualBody)
	}
}
