package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	isQuestion := func(s string) bool {
		return strings.HasSuffix(s, "?")
	}
	isYelled := func(s string) bool {
		return strings.ToUpper(s) == s &&
			strings.ToLower(s) != strings.ToUpper(s)
	}

	s := strings.TrimSpace(remark)
	switch {
	case isQuestion(s) && isYelled(s):
		return "Calm down, I know what I'm doing!"
	case isQuestion(s):
		return "Sure."
	case isYelled(s):
		return "Whoa, chill out!"
	case s == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
