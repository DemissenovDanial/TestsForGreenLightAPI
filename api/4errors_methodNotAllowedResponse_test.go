package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMethodNotAllowedResponse(t *testing.T) {
	app := &application{}

	req := httptest.NewRequest("POST", "/test", nil)
	rr := httptest.NewRecorder()

	app.methodNotAllowedResponse(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d; got %d", http.StatusMethodNotAllowed, rr.Code)
	}

	expectedBody := `{"error":"the POST method is not supported this resource"}`
	actualBody := rr.Body.String()
	actualBody = strings.Replace(actualBody, "\n", "", -1)
	actualBody = strings.Replace(actualBody, "\t", "", -1)
	actualBody = strings.Replace(actualBody, ": ", ":", -1)
	if actualBody != expectedBody {
		t.Errorf("expected body %q; got %q", expectedBody, actualBody)
	}
}
