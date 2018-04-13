// Package acronym exposes abbreviation helpers
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes an arbitrary string and makes it into an acronymn
func Abbreviate(s string) string {
	r := []rune(s)

	// Start with the first letter
	result := string(r[0])

	// len(r)-1 to avoid invalid rune selection at end of string
	for i := 0; i < len(r)-1; i++ {
		firstSelection := r[i]
		secondSelection := r[i+1]

		if unicode.IsLetter(firstSelection) == false && unicode.IsLetter(secondSelection) == true {
			acronymLetter := string(r[i+1])
			result += strings.ToUpper(acronymLetter)
		}
	}
	return result
}
