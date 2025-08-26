package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestApplication_AllUsers(t *testing.T) {
	var mockedRows = mockedDB.NewRows([]string{"id", "email", "first_name", "last_name", "password", "user_active", "created_at", "updated_at", "has_token"})
	mockedRows.AddRow(1, "john@example.com", "John", "Doe", "password", 1, time.Now(), time.Now(), 1)
	mockedRows.AddRow(2, "jane@example.com", "Jane", "Smith", "password", 1, time.Now(), time.Now(), 0)

	mockedDB.ExpectQuery("select id, email, first_name, last_name, password, user_active, created_at, updated_at, case when").WillReturnRows(mockedRows)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/admin/users", nil)
	handler := http.HandlerFunc(testApp.AllUsers)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rr.Code)
	}
}
