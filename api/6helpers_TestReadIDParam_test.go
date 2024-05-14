package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReadIDParam(t *testing.T) {
	req := httptest.NewRequest("GET", "/path/to/resource/123", nil)

	params := httprouter.Params{
		{Key: "id", Value: "123"},
	}

	ctx := context.WithValue(req.Context(), httprouter.ParamsKey, params)

	req = req.WithContext(ctx)

	app := &application{}

	id, err := app.readIDParam(req)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expectedID := int64(123)
	if id != expectedID {
		t.Errorf("got %d, want %d", id, expectedID)
	}
}
