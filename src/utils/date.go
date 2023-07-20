package utils

import (
	"fmt"
	"time"
)

func timeAgo(timestamp time.Time) string {
	now := time.Now()
	duration := now.Sub(timestamp)

	switch {
	case duration >= 365*24*time.Hour:
		years := int64(duration / (365 * 24 * time.Hour))
		return formatTimeAgo(years, "year")
	case duration >= 30*24*time.Hour:
		months := int64(duration / (30 * 24 * time.Hour))
		return formatTimeAgo(months, "month")
	case duration >= 24*time.Hour:
		days := int64(duration / (24 * time.Hour))
		return formatTimeAgo(days, "day")
	case duration >= time.Hour:
		hours := int64(duration / time.Hour)
		return formatTimeAgo(hours, "hour")
	case duration >= time.Minute:
		minutes := int64(duration / time.Minute)
		return formatTimeAgo(minutes, "minute")
	default:
		seconds := int64(duration / time.Second)
		return formatTimeAgo(seconds, "second")
	}
}

func formatTimeAgo(diff int64, unit string) string {
	if diff == 1 {
		return fmt.Sprintf("%d %s ago", diff, unit)
	}
	return fmt.Sprintf("%d %ss ago", diff, unit)
}