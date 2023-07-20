package main

import (
	"overflowing/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", routes.Hello)

	e.GET("/stats", routes.Stats)

	e.Logger.Fatal(e.Start(":4321"))
}
