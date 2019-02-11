package luhn

import (
	"strconv"
	"strings"
)

//Valid calculate the Luhn checksum of a string of digits
func Valid(s string) bool {
	sum := 0
	n := strings.Join(strings.Fields(s), "")
	if len(n) < 2 {
		return false
	}
	even := func(k int) bool { return k%2 == 0 }
	for i := range n {
		x, err := strconv.Atoi(string(n[i]))
		if err != nil {
			return false
		}
		tosum := x
		if (even(len(n)) && even(i)) || (!even(len(n)) && !even(i)) {
			tosum = (x * 2)
			if tosum > 9 {
				tosum -= 9
			}
		}
		sum += tosum
	}
	return sum%10 == 0
}
