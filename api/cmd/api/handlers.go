package main

import (
	"net/http"
	"time"
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

	app.infoLog.Printf("Received login request for user: %s, password: %s", creds.UserName, creds.Password)

	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	payload = jsonResponse{
		Error:   false,
		Message: "Login successful",
		Data: envelope{
			"token": token,
			"user":  user,
		},
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println("Error writing JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorLog.Println("Error reading JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
		app.errorLog.Println("Error deleting token:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Logout successful",
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println("Error writing JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
