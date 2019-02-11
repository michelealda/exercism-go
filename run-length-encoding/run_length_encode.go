package encode

import (
	"strconv"
	"strings"
	"unicode"
)

type item struct {
	k     rune
	count int
}

//RunLengthEncode using RLE
func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return s
	}
	counter := []item{}

	for i, c := range s {
		if i == 0 {
			counter = append(counter, item{c, 1})
		} else if c == rune(s[i-1]) {
			counter[len(counter)-1].count++
		} else {
			counter = append(counter, item{c, 1})
		}

	}

	result := ""
	for _, c := range counter {
		if c.count > 1 {
			result = result + strconv.Itoa(c.count) + string(c.k)
		} else {
			result += string(c.k)
		}

	}
	return result
}

//RunLengthDecode using RLE
func RunLengthDecode(s string) string {
	if len(s) == 0 {
		return s
	}
	result := ""
	for i, c := range s {
		if unicode.IsNumber(c) {
			continue
		}

		if i > 0 {
			p := rune(s[i-1])

			if unicode.IsNumber(p) {
				times, _ := strconv.Atoi(string(p))
				result += strings.Repeat(string(c), times)
			}
		} else {
			result += string(c)
		}

	}
	return result
}
