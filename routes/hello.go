package routes

import (
	"github.com/labstack/echo/v4"
)

func Hello(e echo.Context) error {
	return e.String(200, "Hello, World!")
}
