package routes

import (
	"fmt"
	"overflowing/src/structs"
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

	theme := structs.Theme{"#9745f5", "#9f4bff", "#ffffff", "#000000"};

	svg, err := utils.GenerateSVG(stats, theme);

	e.Set("Content-Type", "image/svg+xml; charset=utf-8")

	return e.String(200, svg)
}