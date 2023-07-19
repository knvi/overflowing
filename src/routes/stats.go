package routes

import (
	"fmt"
	"overflowing/src/utils"

	"github.com/labstack/echo/v4"
)

func Stats(e echo.Context) error {
	userId := e.QueryParam("id")
	fmt.Sprintf("requesting stats for %s", userId)

	stats, err := utils.GetStats(userId)

	if err != nil {
		return e.String(500, err.Error())
	}

	return e.JSON(200, stats)
}