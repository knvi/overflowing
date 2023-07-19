package main

import (
	"overflowing/src/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", routes.Hello)
	e.Logger.Fatal(e.Start(":8080"))
}
