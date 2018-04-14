// Package hamming exports hamming related DNA functions
package hamming

import "errors"

// Distance calculates the hamming difference between two
// distinct DNA strands
func Distance(a, b string) (int, error) {
	var result int

	if len(a) != len(b) {
		return -1, errors.New("strand lengths do not match")
	}

	for i := 0; len(a) > i; i++ {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
