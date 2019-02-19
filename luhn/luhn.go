package luhn

import (
	"strconv"
	"strings"
)

func isSecondDigit(n string, i int) bool {
	return len(n)%2 == 0 && i%2 == 0 ||
		len(n)%2 == 1 && i%2 == 1
}

//Valid calculate the Luhn checksum of a string of digits
func Valid(s string) bool {
	sum := 0
	s = strings.Replace(s, " ", "", -1)
	if len(s) < 2 {
		return false
	}

	for i := range s {
		x, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		tosum := x
		if isSecondDigit(s, i) {
			tosum = (x * 2)
			if tosum > 9 {
				tosum -= 9
			}
		}
		sum += tosum
	}
	return sum%10 == 0
}
