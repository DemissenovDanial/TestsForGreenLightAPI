package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateMovieHandler(t *testing.T) {
	app := &application{}

	var jsonData = []byte(`{"title":"Test Movie","year":int32(2022),"runtime":120,"genres":["Action","Thriller"]}`)

	req, err := http.NewRequest("POST", "/v1/movies", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.createMovieHandler(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusCreated)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("unable to unmarshal response body: %v", err)
	}
	movie, ok := response["movie"].(map[string]interface{})
	if !ok {
		t.Error("response does not contain movie data")
	}
	expectedTitle := "Test Movie"
	if movie["title"] != expectedTitle {
		t.Errorf("unexpected movie title: got %s want %s", movie["title"], expectedTitle)
	}
}
