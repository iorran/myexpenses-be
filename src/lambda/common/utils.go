package common

import "time"

func ValidateAndParseDate(date string) time.Time {
	if date == "" {
		LogError("Date is empty", nil)
	}
	parsedDate, err := time.Parse(time.DateOnly, date)

	LogError("Error when parsing string to date", err)
	return parsedDate
}
