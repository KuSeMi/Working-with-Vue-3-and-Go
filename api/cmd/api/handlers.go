package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println("Error reading JSON:", err)
		payload.Error = true
		payload.Message = "Invalid request payload"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// TODO authenicate
	app.infoLog.Printf("Received login request for user: %s, password: %s", creds.UserName, creds.Password)

	// send back a JSON response
	payload.Error = false
	payload.Message = "Login successful"
	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println("Error writing JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
