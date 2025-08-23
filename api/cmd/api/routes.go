package main

import (
	"api/internal/data"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Accept", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/users/login", app.Login)
	mux.Post("/users/logout", app.Logout)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
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

	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "Poroh@ua.com",
			Password:  "password",
			FirstName: "Vadym",
			LastName:  "Krasavets",
		}

		app.infoLog.Println("Adding user:", u)

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("New user ID:", id)
		newUser, _ := app.models.User.GetOne(id)
		app.writeJSON(w, http.StatusCreated, newUser)
	})

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.Email = "admin@i.ua"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "Token generated successfully",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.Token.GenerateToken(2, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		user, err := app.models.User.GetOne(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Token saved successfully",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err, http.StatusUnauthorized)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "Token validation successful",
			Data:    valid,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	return mux
}
