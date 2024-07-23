// Package gigasecond is a way of working with 1 thousand million seconds
package gigasecond

import "time"

// AddGigasecond determines the date and time one gigasecond after a given date
func AddGigasecond(t time.Time) time.Time {
	giga := 1_000_000_000
	return t.Add(time.Second * time.Duration(giga))
}
