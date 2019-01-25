package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	items := regexp.
		MustCompile(`[\w']+`).
		FindAllString(s, -1)
	runes := []rune{}
	for _, s := range items {
		runes = append(runes, []rune(s)[0])
	}
	return strings.ToUpper(string(runes))
}
