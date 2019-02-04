package isogram

import (
	"regexp"
	"strings"
)

//IsIsogram return if a word has no repeating characters
func IsIsogram(word string) bool {
	w := strings.ToLower(
		strings.Join(
			regexp.MustCompile(`\w+`).FindAllString(word, -1),
			""))

	runeCount := map[rune]int{}
	for _, r := range w {
		_, ok := runeCount[r]
		if ok {
			return false
		} else {
			runeCount[r] = 1
		}
	}
	return true
}
