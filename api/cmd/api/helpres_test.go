package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_readJSON(t *testing.T) {
	sampleJSON := map[string]interface{}{
		"id":         1,
		"email":      "john@example.com",
		"first_name": "John",
		"last_name":  "Doe",
		"password":   "password",
		"created_at": time.Now(),
		"updated_at": time.Now(),
		"has_token":  1,
	}

	body, _ := json.Marshal(sampleJSON)
	var decodedJSON struct {
		ID        int       `json:"id"`
		Email     string    `json:"email"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		HasToken  int       `json:"has_token"`
	}

	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
		return
	}

	rr := httptest.NewRecorder()
	defer req.Body.Close()

	err = testApp.readJSON(rr, req, &decodedJSON)
	if err != nil {
		t.Error("Failed to decode JSON", err)
	}
}

func Test_writeJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := jsonResponse{
		Error:   false,
		Message: "it worked",
	}

	headers := make(http.Header)
	headers.Add("Location", "/somewhere")

	err := testApp.writeJSON(rr, http.StatusCreated, payload, headers)
	if err != nil {
		t.Error("Failed to write JSON", err)
	}

	testApp.environment = "production"
	err = testApp.writeJSON(rr, http.StatusCreated, payload, headers)
	if err != nil {
		t.Error("Failed to write JSON in production environment:", err)
	}

	testApp.environment = "development"
}

func Test_errorJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	err := testApp.errorJSON(rr, errors.New("some error"))
	if err != nil {
		t.Error(err)
	}

	testJSONPayload(t, rr)

	errSSlice := []string{
		"(SQLSTATE 23505)",
		"(SQLSTATE 22001)",
		"(SQLSTATE 23503)",
	}

	for _, sqlError := range errSSlice {
		customErr := testApp.errorJSON(rr, errors.New(sqlError), http.StatusUnauthorized)
		if customErr != nil {
			t.Error(customErr)
		}
		testJSONPayload(t, rr)
	}
}

func testJSONPayload(t *testing.T, rr *httptest.ResponseRecorder) {
	var requestPayload jsonResponse
	if err := json.NewDecoder(rr.Body).Decode(&requestPayload); err != nil {
		t.Error("received error when decoding errorJSON payload", err)
	}

	if !requestPayload.Error {
		t.Error("errorJSON returned false for error field when it should be true")
	}
}
