package scrabble

import "strings"

//Score will calculate the scrabble score of a word
func Score(input string) int {
	runeValue := func(r rune) int {
		switch {
		case strings.ContainsRune("DG", r):
			return 2
		case strings.ContainsRune("BCMP", r):
			return 3
		case strings.ContainsRune("FHVWY", r):
			return 4
		case strings.ContainsRune("K", r):
			return 5
		case strings.ContainsRune("JX", r):
			return 8
		case strings.ContainsRune("QZ", r):
			return 10
		default:
			return 1
		}
	}

	score := 0
	for _, r := range strings.ToUpper(input) {
		score += runeValue(r)
	}
	return score
}
