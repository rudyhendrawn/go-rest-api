package main

import (
	"go-echo-app/db"
	"go-echo-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
