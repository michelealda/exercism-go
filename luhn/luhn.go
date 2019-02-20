package luhn

import (
	"strconv"
	"strings"
)

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
		if len(s)%2 == i%2 {
			tosum = (x * 2)
			if tosum > 9 {
				tosum -= 9
			}
		}
		sum += tosum
	}
	return sum%10 == 0
}
