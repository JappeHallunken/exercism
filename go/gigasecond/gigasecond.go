// Package gigasecond provides a function to add a gigasecond to a given time.
// A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

import "time"

// AddGigasecond takes a time.Time value and returns a new time.Time value
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
