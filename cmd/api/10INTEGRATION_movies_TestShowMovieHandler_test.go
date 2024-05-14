package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowMovieHandler(t *testing.T) {
	app := &application{}

	movieID := 123
	req := httptest.NewRequest("GET", fmt.Sprintf("/v1/movies/%d", movieID), nil)

	rr := httptest.NewRecorder()

	app.showMovieHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("unable to unmarshal response body: %v", err)
	}
	movie, ok := response["movie"].(map[string]interface{})
	if !ok {
		t.Error("response does not contain movie data")
	}
	expectedID := float64(movieID)
	if movie["id"] != expectedID {
		t.Errorf("unexpected movie ID: got %v want %v", movie["id"], expectedID)
	}
}
