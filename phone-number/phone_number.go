package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

//Number create a NANP from a string
func Number(s string) (string, error) {
	x := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, s)

	if x[0] == '1' {
		x = x[1:]
	}
	if len(x) != 10 {
		return "", errors.New("invalid lenght")
	}
	if x[0] < '2' || x[3] < '2' {
		return "", errors.New("invalid digits")
	}

	return x, nil
}

//AreaCode return the area code of a NANP
func AreaCode(s string) (string, error) {
	x, e := Number(s)
	if e != nil {
		return "", e
	}
	return x[:3], nil
}

//Format print a number in a nice format
func Format(s string) (string, error) {
	x, e := Number(s)
	if e != nil {
		return "", e
	}
	return fmt.Sprintf("(%s) %s-%s", x[:3], x[3:6], x[6:]), nil
}
