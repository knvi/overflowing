package utils

import (
	"overflowing/src/structs"
	"strconv"
)

func GenerateSVG(stats structs.Stats, theme structs.Theme) (svg string, error error) {
	const height = 47
	const fontSize = 12
	const badgeGap = 18
	const oPadding = 12
	const iPadding = 6
	const imgSize = 24

	// calculate the width of the svg
	width := calculateWidth(stats.Reputation, stats.Gold, stats.Silver, stats.Bronze, badgeGap)

	svg += `<svg data-userId="` + stats.ID + `" width="` + str(width) + `" height="` + str(height) + `" viewBox="0 0 ` + str(width) + ` ` + str(height) + `" fill="none" xmlns="http://www.w3.org/2000/svg">`
	svg += `<rect width="` + str(width) + `" height="` + str(height) + `" fill="` + theme.BgColor + `"/>`

	if stats.ImageUrl != "" {
		img, err := ImageToBase64(stats.ImageUrl)
		if err != nil {
			return "", err
		}
		svg += genImage(img, oPadding, imgSize)
	}

	svg += genRep(oPadding+imgSize+iPadding+18, stats, theme, height, fontSize)

	badgeXPos := oPadding + imgSize + iPadding*2 + calcRepWidth(stats.Reputation) + iPadding*2
	if stats.Gold > 0 {
		svg += genBadge("Gold", badgeXPos, height/2, stats.Gold, fontSize, theme.Gold)
	}

	if stats.Silver > 0 {
		badgeXPos = badgeXPos + calcBadgeGap(stats.Gold) + badgeGap
		svg += genBadge("Silver", badgeXPos, height/2, stats.Silver, fontSize, theme.Silver)
	}

	if stats.Bronze > 0 {
		badgeXPos = badgeXPos + calcBadgeGap(stats.Silver) + badgeGap
		svg += genBadge("Bronze", badgeXPos, height/2, stats.Bronze, fontSize, theme.Bronze)
	}

	svg += `</svg>`

	return svg, nil
}

func genImage(imageBase64 string, xPos, size int) (svg string) {
	fullImage := "data:image/png;base64," + imageBase64
	svg = ` <image x=" ` + str(xPos) + `" y="10" href="` + fullImage + `" height="` + str(size) + `" width="` + str(size) + `" default-src="sha256-4Su6mBWzEIFnH4pAGMOuaeBrstwJN4Z3pq/s1Kn4/KQ=" />`

	return svg
}

func genBadge(id string, xPos, yPos, count, fontSize int, color string) (svg string) {
	gap := calcBadgeGap(count)
	const radius = 3

	svg += `<circle text-anchor="middle" dominant-baseline="middle" cx="` + str(xPos) + `" cy="` + str(yPos) + `" r="` + str(radius) + `" fill="` + color + `"/>`
	svg += `<text data-testBadge` + id + `="` + str(count) + `" x="` + str(xPos+gap) + `" y="` + str(yPos) + `" font-size="` + str(fontSize) + `" font-family="Arial" font-weight="bold" text-anchor="middle" dominant-baseline="middle" fill="` + color + `">` + str(count) + `</text>`

	return svg
}

func genRep(xPos int, stats structs.Stats, theme structs.Theme, height int, fontSize int) string {
	svg := `<text data-testReputation="` + str(stats.Reputation) + `"  x="` + str(xPos) + `" y="` + str(height/2) + `" font-weight="bold" fill="` + theme.TextColor + `" font-family="Arial" font-size="` + str(fontSize) + `" text-anchor="middle" dominant-baseline="middle">` + str(stats.Reputation) + `</text>`
	return svg
}

func calculateWidth(rep, gold, silver, bronze, badgeGap int) int {
	min := 83 // minimal width
	scaler := 0.5 // how much to scale the width

	for _, v := range []int{gold, silver, bronze} {
		if v > 0 {
			min += calcBadgeGap(v) + scale(calcBadgeGap(v), scaler) + scale(badgeGap, scaler)
		}
	}

	width := min + calcRepWidth(rep)

	return width
}

func scale(v int, factor float64) int {
	return int(float64(v) * factor)
}

func calcBadgeGap(count int) int {
	if count == 0 {
		return 0
	} else if count < 10 {
		return 9
	} else if count < 100 {
		return 14
	} else if count < 1000 {
		return 16
	}
	return 18
}

func calcRepWidth(rep int) int {
	if rep < 10 {
		return 7
	} else if rep < 100 {
		return 12
	} else if rep < 1000 {
		return 20
	}
	return 33
}

func str(n int) string {
	return strconv.Itoa(n)
}