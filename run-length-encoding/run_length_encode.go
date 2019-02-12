package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//RunLengthEncode using RLE
func RunLengthEncode(s string) string {
	result := ""
	c := 1
	for i, x := range s {
		if i == 0 {
			continue
		}
		if x == rune(s[i-1]) {
			c++
		} else {
			result, c = add(c, i, s, result)
		}
	}
	// last item is counted but not yet added
	if len(s) > 0 {
		result, _ = add(c, len(s), s, result)
	}
	return result
}

func add(times, index int, original, result string) (string, int) {
	if times > 1 {
		result += fmt.Sprintf("%d%c", times, original[index-1])
		times = 1
	} else {
		result += fmt.Sprintf("%c", original[index-1])
	}
	return result, times
}

//RunLengthDecode using RLE
func RunLengthDecode(s string) string {
	result, countAsString := "", ""
	for _, x := range s {
		if unicode.IsDigit(x) {
			countAsString += string(x)
		} else {
			c, err := strconv.Atoi(countAsString)
			if err != nil {
				result += string(x)
			} else {
				result += strings.Repeat(string(x), c)
			}
			countAsString = ""
		}
	}
	return result
}
