package helpers

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func TimeToPointer(data time.Time) *time.Time {
	return &data
}

func DurationToString(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	var parts []string

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}

func GetDurations(startAt time.Time, interval time.Duration) []time.Time {
	var arr []time.Time

	date := startAt.Round(interval)

	for {
		if date.Unix() >= time.Now().Unix() {
			return arr
		}

		arr = append(arr, date)

		date = date.Add(interval)
	}

	return arr
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Truncate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}
