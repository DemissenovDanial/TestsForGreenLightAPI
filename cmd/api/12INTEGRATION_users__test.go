package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterUserHandler(t *testing.T) {
	app := &application{}

	newUser := map[string]interface{}{
		"name":     "John Doe",
		"email":    "john@example.com",
		"password": "secretpassword",
	}
	jsonBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("failed to marshal JSON body: %v", err)
	}

	req, err := http.NewRequest("POST", "/v1/users", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	app.registerUserHandler(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusAccepted)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("unable to unmarshal response body: %v", err)
	}
	user, ok := response["user"].(map[string]interface{})
	if !ok {
		t.Error("response does not contain user data")
	}
	expectedName := "John Doe"
	if user["name"] != expectedName {
		t.Errorf("unexpected user name: got %v want %v", user["name"], expectedName)
	}
}
