// Package raindrops exposes raindrop-related functions
package raindrops

import "strconv"

// Convert takes a number and converts it into the beautiful sound of raindrops.
func Convert(s int) (result string) {
	if s%3 == 0 {
		result += "Pling"
	}
	if s%5 == 0 {
		result += "Plang"
	}
	if s%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(s)
	}
	return
}
