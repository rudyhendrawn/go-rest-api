package handlers

import (
	"net/http"
	"strconv"

	"go-echo-app/db"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := db.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get user"})
	}

	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	users, err := db.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to get user"})
	}

	return c.JSON(http.StatusOK, users)
}
