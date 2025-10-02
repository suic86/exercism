package booking

import (
	"time"
)

const (
	scheduleLayout               string = "1/2/2006 15:04:05"
	hasPassedLayout              string = "January 2, 2006 15:04:05"
	isAfternoonAppointmentLayout string = "Monday, January 2, 2006 15:04:05"
	descriptionParseLayout       string = "1/2/2006 15:04:05"
	descriptionFormatLayout      string = "You have an appointment on Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse(scheduleLayout, date)
	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse(hasPassedLayout, date)
	if err != nil {
		panic(err)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse(isAfternoonAppointmentLayout, date)
	if err != nil {
		panic(err)
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, err := time.Parse(descriptionParseLayout, date)
	if err != nil {
		panic(err)
	}
	return t.Format(descriptionFormatLayout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
