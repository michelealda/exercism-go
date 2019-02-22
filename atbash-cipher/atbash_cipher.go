package atbash

import (
	"unicode"
)

const cipher = "zyxwvutsrqponmlkjihgfedcba"

func encode(r rune) rune {
	if unicode.IsDigit(r) {
		return r
	}
	if !unicode.IsLetter(r) {
		return -1
	}
	//rune('a') = 97
	return rune(cipher[unicode.ToLower(r)-97])
}

//Atbash encode a string using the Atbash chiper
func Atbash(s string) string {
	result := ""
	count := 0
	for i, r := range s {
		e := encode(r)
		if e < 0 {
			continue
		}
		if count%5 == 0 && i > 0 {
			result += " "
		}
		count++
		result += string(e)
	}

	return result
}
