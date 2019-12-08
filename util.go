package gowizz

import (
	"time"
)

const (
	// Day Duration of a single day
	Day = 24 * time.Hour
)

// GenTimeRanges Generates time ranges
func GenTimeRanges(start time.Time, duration time.Duration, count int) []TimeRange {
	result := make([]TimeRange, count)
	startDate := start.Truncate(24 * time.Hour)

	for i := 0; i < count; i++ {
		endDate := startDate.Add(duration)
		result[i] = TimeRange{
			From: startDate,
			To:   endDate,
		}

		startDate = endDate.Add(Day)
	}

	return result
}

type TimeRange struct {
	From time.Time
	To   time.Time
}
