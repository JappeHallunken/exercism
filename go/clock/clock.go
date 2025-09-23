package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour, minute int
}

func New(h, m int) Clock {
	total := normalize(h*60 + m)
	return Clock{hour: total / 60, minute: total % 60}
}

func (c Clock) Add(m int) Clock {
	total := normalize(c.hour*60 + c.minute + m)
	return Clock{hour: total / 60, minute: total % 60}
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func normalize(total int) int {
	const minutesPerDay = 24 * 60
	total %= minutesPerDay
	if total < 0 {
		total += minutesPerDay
	}
	return total
}
