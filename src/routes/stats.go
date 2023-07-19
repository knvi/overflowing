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

	theme := structs.Theme{Gold: "#F0B400", Silver: "#999C9F", Bronze: "#AB8A5F", BgColor: "#2D2D2D", TextColor: "#C4CCBC"}

	svg, err := utils.GenerateSVG(stats, theme);

	e.Set("Content-Type", "image/svg+xml; charset=utf-8")

	return e.String(200, svg)
}	