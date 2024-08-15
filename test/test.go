package test

import (
	"den/main"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupApp() *main.App {
	app := main.App()
	// Inisialisasi aplikasi Anda di sini
	// Misalnya, jika Anda menggunakan framework seperti Fiber atau Echo, Anda bisa menginisialisasinya di sini
	return &app{}
}

func TestGetUsers(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest("GET", "/users", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateUser(t *testing.T) {
	app := setupApp()

	payload := `{"name": "John Doe", "email": "john@example.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdateUser(t *testing.T) {
	app := setupApp()

	// Create a user first
	payload := `{"name": "John Doe", "email": "john@example.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	// Update the user
	updatePayload := `{"name": "John Smith", "email": "johnsmith@example.com"}`
	req = httptest.NewRequest("PUT", "/users/1", strings.NewReader(updatePayload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteUser(t *testing.T) {
	app := setupApp()

	// Create a user first
	payload := `{"name": "John Doe", "email": "john@example.com"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	app.Test(req)

	// Delete the user
	req = httptest.NewRequest("DELETE", "/users/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
