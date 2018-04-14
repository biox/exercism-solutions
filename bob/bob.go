// Package bob exports functions related to our apathetic teenage friend.
package bob

import (
	"strings"
)

func yelly(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}

func questiony(s string) bool {
	return strings.HasSuffix(s, "?")
}

// Hey is bob's personality function. "Siri for teenagers"
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case yelly(remark) && questiony(remark):
		return "Calm down, I know what I'm doing!"
	case yelly(remark):
		return "Whoa, chill out!"
	case questiony(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}
