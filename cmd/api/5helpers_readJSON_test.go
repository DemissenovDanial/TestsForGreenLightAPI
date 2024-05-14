package main

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestReadJSON(t *testing.T) {
	app := &application{}

	jsonData := `{"name": "John", "age": 30}`
	req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(jsonData))
	rr := httptest.NewRecorder()

	var data struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err := app.readJSON(rr, req, &data)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expectedName := "John"
	expectedAge := 30
	if data.Name != expectedName || data.Age != expectedAge {
		t.Errorf("expected name=%q, age=%d; got name=%q, age=%d", expectedName, expectedAge, data.Name, data.Age)
	}
}
