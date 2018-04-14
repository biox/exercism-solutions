// Package twofer exposes some trashy string concat functions.
package twofer

// ShareWith shares with formally you, or possibly just you.
func ShareWith(s string) string {
	if s == "" {
		s = "you"
	}

	return "One for " + s + ", one for me."
}
