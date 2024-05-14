package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateMovieHandler(t *testing.T) {
	app := &application{}

	existingMovieID := 123
	updatedMovieData := map[string]interface{}{
		"title":   "Updated Movie Title",
		"year":    2023,
		"runtime": 150,
		"genres":  []string{"Drama", "Romance"},
	}
	updatedJSONData, err := json.Marshal(updatedMovieData)
	if err != nil {
		t.Fatalf("failed to marshal updated movie data: %s", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("/v1/movies/%d", existingMovieID), bytes.NewBuffer(updatedJSONData))
	if err != nil {
		t.Fatalf("failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.updateMovieHandler(rr, req)

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
	expectedTitle := "Updated Movie Title"
	if movie["title"] != expectedTitle {
		t.Errorf("unexpected movie title: got %v want %v", movie["title"], expectedTitle)
	}
}
