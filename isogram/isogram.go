// Package isogram exposes functions used to determine the isogramminess of strings
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines whether or not a string is an isogram
func IsIsogram(isogram string) bool {
	var remainder int
	isogram = strings.ToUpper(isogram)

	for _, c := range isogram {
		// Isograms only care about letters
		if !unicode.IsLetter(c) {
			continue
		}

		char := strings.ToUpper(string(c))
		remainder = strings.Count(isogram, char)

		if remainder > 1 {
			return false
		}

		remainder = 0
	}

	return true
}
