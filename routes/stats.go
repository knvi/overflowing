package routes

import (
	"fmt"
	"overflowing/structs"
	"overflowing/utils"

	"github.com/labstack/echo/v4"
)

func Stats(e echo.Context) error {
	userId := e.QueryParam("id")

	if userId == "" {
		return e.String(400, "missing id")
	}

	fmt.Println(fmt.Sprintf("requesting stats for %s", userId))

	stats, err := utils.GetStats(userId)

	if err != nil {
		return e.String(500, err.Error())
	}

	theme := structs.Theme{Gold: "#F0B400", Silver: "#999C9F", Bronze: "#AB8A5F", BgColor: "#2D2D2D", TextColor: "#C4CCBC"}

	svg, err := utils.GenerateSVG(stats, theme);

	if err != nil {
		return e.String(500, err.Error())
	}

	cache := 3600 * 24

	// set headers

	e.Response().Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	e.Response().Header().Set("Cache-Control", "public, max-age="+ fmt.Sprintf("%d", cache))
	e.Response().Header().Set("ContentSecurityPolicy", "default-src 'none'; style-src 'unsafe-inline'; img-src data:;")
	return e.String(200, svg)
}	