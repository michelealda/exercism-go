package pangram

import "strings"

const letters = "abcdefghijklmnopqrstuvwxyz"

//IsPangram determines if a sentence is a pangram
func IsPangram(input string) bool {
	lc := strings.ToLower(input)
	for _, r := range letters {
		if strings.IndexRune(lc, r) < 0 {
			return false
		}
	}
	return true
}
