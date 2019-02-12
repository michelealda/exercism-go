package isbn

import "strconv"

//IsValidISBN check the validity of a ISBN number
func IsValidISBN(s string) bool {
	i, sum := 10, 0
	for _, r := range s {
		if r == '-' {
			continue
		}
		if x, err := strconv.Atoi(string(r)); err == nil {
			sum += x * i
			i--
		} else {
			if i == 1 && r == 'X' {
				sum += 10 * i
				i--
			} else {
				return false
			}
		}
	}
	return i == 0 && sum%11 == 0
}
