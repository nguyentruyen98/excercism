package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(dates string) time.Time {
	layout := "1/02/2006 15:04:05"

	t, err := time.Parse(layout, dates)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	current := time.Now()
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	return current.After(parsedTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	if err != nil {
		panic(err)
	}
	if parsedTime.Hour() >= 12 && parsedTime.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)

	formattedTime := parsedTime.Format("Monday, January 2, 2006, at 15:04")

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
