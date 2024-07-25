package util

import (
	"errors"
	"time"
)

func ValidateTime(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date string is empty")
	}
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}
	return parsedDate, nil
}
