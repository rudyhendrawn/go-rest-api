package main

import (
	"context"
	"encoding/json"
	"go-echo-app/db"
	"go-echo-app/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

func TestMain(t *testing.T) {
	// Initialize the database
	db.InitDB()

	// Create a new Echo instance
	e := echo.New()

	// Setup the routes
	// routes.SetupRoutes(e)

	// Start the server in a separate goroutine
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Server is assumed to be started now; you can add additional checks here if needed

	// Stop the server
	if err := e.Shutdown(context.Background()); err != nil {
		t.Fatalf("Failed to shut down server: %v", err)
	}
}

func TestGetAllUsers(t *testing.T) {
	// Initialize the database (consider using a mock or test database)
	db.InitDB()

	// Create a new Echo instance
	e := echo.New()

	// Setup the routes
	// routes.SetupRoutes(e)

	// Create a test server
	ts := httptest.NewServer(e)
	defer ts.Close()

	// Send a request to the "get all users" endpoint
	res, err := http.Get(ts.URL + "/users")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", res.Status)
	}

	// Read and parse the response body
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}

	var users []*models.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Here you can add more assertions
	// For example, check the length of the users slice
	// or verify the content of the returned users
}
