package utils

import (
	"fmt"
)

func FormatNumber(num int) string {
	if num < 10000 {
		return fmt.Sprintf("%d", num)
	} else if num < 1000000 {
		return fmt.Sprintf("%.1fk", float64(num)/1000)
	} else if num < 1000000000 {
		return fmt.Sprintf("%.1fM", float64(num)/1000000)
	} else if num < 1000000000000 {
		return fmt.Sprintf("%.1fB", float64(num)/1000000000)
	} else {
		return fmt.Sprintf("%.1fT", float64(num)/1000000000000)
	}
}