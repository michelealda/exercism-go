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
	parity := len(s) % 2
	for i := range s {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
