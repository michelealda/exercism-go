package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(y int) bool {
	return y%400 == 0 || y%100 != 0 && y%4 == 0
}
