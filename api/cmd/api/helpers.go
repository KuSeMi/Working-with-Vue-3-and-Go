package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(data); err != nil {
		return err
	}

	err := decoder.Decode(&struct{}{}) // ensure no trailing data
	if err != io.EOF {
		return errors.New("request body must only contain a single JSON object")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	if _, err := w.Write(out); err != nil {
		return err
	}

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var customError error

	switch {
	case strings.Contains(err.Error(), "SQLSTATE 23505"):
		customError = errors.New("duplicate key value violates unique constraint")
		statusCode = http.StatusForbidden
	case strings.Contains(err.Error(), "SQLSTATE 22001"):
		customError = errors.New("value too long for column")
		statusCode = http.StatusBadRequest
	case strings.Contains(err.Error(), "SQLSTATE 23503"):
		customError = errors.New("foreign key violation")
		statusCode = http.StatusForbidden
	default:
		customError = err
	}

	payload := jsonResponse{
		Error:   true,
		Message: customError.Error(),
	}

	app.writeJSON(w, statusCode, payload)
}
