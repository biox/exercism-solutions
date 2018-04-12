// Package gigasecond exposes functions related to changing time.
// Time values in orders of gigaseconds.
package gigasecond

import (
	"time"
)

// AddGigasecond simply adds one gigasecond (10^9) to a time.
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Duration(expo(10, 9))
	result := t.Add(time.Second * gigasecond)
	return result
}

// Whatever I'll do it myself goddammit!
func expo(a, b int) int {
	result := 1

	for b > 0 {
		result *= a
		b--
	}

	return result
}
