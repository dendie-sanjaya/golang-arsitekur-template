package main

import (
	"bytes"
	"den/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// setupMockDB creates a mock database connection using GORM and sqlmock
func setupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestPostUser(t *testing.T) {
	db, _, err := setupMockDB()
	if err != nil {
		t.Fatalf("Failed to setup mock database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	app := fiber.New()
	controller.UserControler(app, db)

	payload := `{"name": "John Doe", "email": "john@example.com"}`
	req := httptest.NewRequest("POST", "/users", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
