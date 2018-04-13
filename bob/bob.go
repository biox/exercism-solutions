// Package bob exports functions related to our apathetic teenage friend.
package bob

import (
	"regexp"
)

// Hey is bob's personality function. "Siri for teenagers"
func Hey(remark string) string {
	var response string

	yellingQuestion, _ := regexp.Compile(`[A-Z]+\?`)
	question, _ := regexp.Compile(`\?(\s+)?$`)

	// if every character in the string is uppercase and there is a question mark at the end
	if isYelling(remark) && yellingQuestion.MatchString(remark) {
		response = "Calm down, I know what I'm doing!"
		// if every character in the string is uppercase
	} else if isYelling(remark) {
		response = "Whoa, chill out!"
		// if the string contains a question mark at the end
	} else if question.MatchString(remark) {
		response = "Sure."
	} else if nothing(remark) {
		response = "Fine. Be that way!"
	} else {
		response = "Whatever."
	}

	return response
}

// IS THIS YELLING? TRUE.
func isYelling(s string) bool {
	onlyUpper, _ := regexp.Compile("[^A-Z]+")
	upperAndLower, _ := regexp.Compile("[^A-Za-z]+")

	upperString := onlyUpper.ReplaceAllString(s, "")
	upperLowerString := upperAndLower.ReplaceAllString(s, "")

	// No characters? No yelling.
	if upperLowerString == "" {
		return false
	}

	return upperString == upperLowerString
}

// Simply return true if the string is empty
func nothing(s string) bool {
	upperAndLower, _ := regexp.Compile(`[^\s]+`)

	nothingString := upperAndLower.ReplaceAllString(s, "")

	return s == nothingString
}
