package main

import (
	"api/internal/data"
	"api/internal/driver"
	"fmt"
	"log"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	config      config
	infoLog     *log.Logger
	errorLog    *log.Logger
	models      data.Models
	environment string
}

func main() {
	var cfg config
	cfg.port = 8081

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := os.Getenv("DSN")
	environment := os.Getenv("ENV")

	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	app := &application{
		config:      cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		models:      data.New(db.SQL),
		environment: environment,
	}

	if err := app.serve(); err != nil {
		app.errorLog.Fatal(err)
	}
}

func (app *application) serve() error {
	app.infoLog.Println("Starting application...")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	app.infoLog.Println("Starting server on port", app.config.port)

	return srv.ListenAndServe()
}
