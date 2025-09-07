package booking

import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {

	layout := "1/02/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	l := "January 2, 2006 15:04:05"
	t, err := time.Parse(l, date)

	if err != nil {
		fmt.Println(err)
	}

	return time.Since(t) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	l := "Monday, January 2, 2006 15:04:05"
	t, err := time.Parse(l, date)
	if err != nil {
		fmt.Println(err)
	}
	return (t.Hour() >= 12 && t.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var description string
	l := "1/2/2006 15:04:05"
	t, err := time.Parse(l, date)
	if err != nil {
		fmt.Println(err)
	}

	description = fmt.Sprintf("You have an appointment on %v, %v %d, %d, at %d:%d.", t.Weekday().String(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
  yearNow := time.Now().Year()
  month := time.September
  day := 15

  t := time.Date(yearNow, month, day, 0, 0, 0, 0, time.UTC)
  return t
}
