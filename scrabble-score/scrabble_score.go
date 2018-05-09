// Package scrabble does stuff
package scrabble

import (
	"fmt"
	"regexp"
)

// Score takes a scrabble word and returns the point value
func Score(word string) int {
	var tally int

	for _, r := range word {
		tally += derivePoints(r)
	}

	return tally
}

// dervicePoints takes a single rune and derives the point-worth based on some pre-defined regexes
func derivePoints(r rune) (value int) {
	letter := string(r)

	onePoint, _ := regexp.MatchString("[AEIOULNRSTaeioulnrst]", letter)
	fmt.Println(onePoint)
	if onePoint == true {
		return 1
	}

	twoPoints, _ := regexp.MatchString("[DdGg]", letter)
	if twoPoints {
		return 2
	}

	threePoints, _ := regexp.MatchString("[BbCcMmPp]", letter)
	if threePoints {
		return 3
	}

	fourPoints, _ := regexp.MatchString("[FfHhVvWwYy]", letter)
	if fourPoints {
		return 4
	}

	fivePoints, _ := regexp.MatchString("[Kk]", letter)
	if fivePoints {
		return 5
	}

	eightPoints, _ := regexp.MatchString("[JjXx]", letter)
	if eightPoints {
		return 8
	}

	tenPoints, _ := regexp.MatchString("[QqZz]", letter)
	if tenPoints {
		return 10
	}

	return
}
