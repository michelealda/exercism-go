package isogram

import (
	"unicode"
)

//IsIsogram return if a word has no repeating characters
func IsIsogram(word string) bool {
	seen := map[rune]bool{}
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		lcr := unicode.ToLower(r)

		if seen[lcr] {
			return false
		}
		seen[lcr] = true
	}
	return true
}
