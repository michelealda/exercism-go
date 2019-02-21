package cryptosquare

import (
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

func findCoulumnSize(n int) int {
	for r := 1.0; ; r++ {
		c := math.Ceil(float64(n) / r)
		if c >= r && c-r <= 1.0 {
			return int(c)
		}
	}
}

//Encode a string using a cryptosquare
func Encode(s string) string {
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)

	if len(s) <= 1 {
		return s
	}

	c := findCoulumnSize(len(s))

	xs := []string{}
	for {
		if len(s) <= c {
			for i := utf8.RuneCountInString(s); i < c; i++ {
				s += " "
			}
			xs = append(xs, s)
			break
		}
		xs = append(xs, s[:c])
		s = s[c:]
	}

	rs := ""
	for i := 0; i < c; i++ {
		for j := 0; j < len(xs); j++ {
			rs += string(xs[j][i])
		}
		if i < c-1 {
			rs += " "
		}
	}

	return rs
}
