package util

import "time"

func ValidateTime(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, nil
	}
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
