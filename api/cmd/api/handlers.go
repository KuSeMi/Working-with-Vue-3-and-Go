package main

import (
	"api/internal/data"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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

	// make sure if user is active
	if user.Active == 0 {
		app.errorJSON(w, errors.New("user is not active"), http.StatusUnauthorized)
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

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	var users data.User
	all, err := users.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Users retrieved successfully",
		Data:    envelope{"users": all},
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) EditUser(w http.ResponseWriter, r *http.Request) {
	var user data.User

	err := app.readJSON(w, r, &user)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		if _, err := app.models.User.Insert(user); err != nil {
			app.errorJSON(w, err, http.StatusBadRequest)
			return
		}
	} else {
		u, err := app.models.User.GetOne(user.ID)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		u.Email = user.Email
		u.FirstName = user.FirstName
		u.LastName = user.LastName
		u.Active = user.Active

		if err := u.Update(); err != nil {
			app.errorJSON(w, err)
			return
		}

		if user.Password != "" {
			err := u.ResetPassword(user.Password)
			if err != nil {
				app.errorJSON(w, err)
				return
			}
		}
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Users changes saved successfully",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid user ID"))
		return
	}

	user, err := app.models.User.GetOne(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, user)
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.models.User.DeleteByID(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "User deleted successfully",
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) LogUserOutAndSetInactive(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user, err := app.models.User.GetOne(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	user.Active = 0
	err = user.Update()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// delete tokens for user
	err = app.models.Token.DeleteTokensForUser(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "user logged out and set to inactive",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *application) ValidateToken(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	valid := false
	valid, _ = app.models.Token.ValidToken(requestPayload.Token)

	payload := jsonResponse{
		Error: false,
		Data:  envelope{"valid": valid},
	}

	app.writeJSON(w, http.StatusOK, payload)
}
