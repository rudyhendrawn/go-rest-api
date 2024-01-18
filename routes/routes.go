package routes

import (
	"go-echo-app/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for User
	e.GET("/users", handlers.GetAllUsers)
	e.GET("/users:id", handlers.GetUser)

	// Test the API
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the API.")
	})
}
