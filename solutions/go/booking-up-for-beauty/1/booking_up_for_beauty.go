package booking

import (
    "time"
    "fmt"
)

//7/25/2019 13:45:00
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Now()
	}
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    parsedTime, err := time.Parse(layout, date)
    if err != nil {
        return false
    }
    return parsedTime.Before(time.Now())
}

// Thursday, July 25, 2019 13:45:00
// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    parsedTime, err := time.Parse(layout, date)
    if err!=nil{
    	return false
    }
    hour := parsedTime.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedTime, err := time.Parse(layout, date)
	if err != nil {
		return "ERROR"
	}

	month := parsedTime.Month()
	hour := parsedTime.Hour()
	minutes := parsedTime.Minute()
	day := parsedTime.Day()
	week := parsedTime.Weekday().String()
	year := parsedTime.Year()
    
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", week, month, day, year, hour, minutes)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
