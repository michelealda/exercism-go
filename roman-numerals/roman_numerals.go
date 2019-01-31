package romannumerals

import (
	"errors"
	"strings"
)

func convert(n, divisor int, single, middle, top string) (string, int) {
	q, m := n/divisor, n%divisor
	result := ""
	switch q {
	case 1, 2, 3:
		result += strings.Repeat(single, q)
	case 4:
		result += single + middle
	case 5:
		result += middle
	case 6, 7, 8:
		result += middle + strings.Repeat(single, q-5)
	case 9:
		result += single + top
	}
	return result, m
}

//ToRomanNumeral converts an arabic digit between 1 and 3000
func ToRomanNumeral(arabic int) (string, error) {
	if arabic > 3000 || arabic < 1 {
		return "", errors.New("out of range input")
	}

	convert1000, hundreds := convert(arabic, 1000, "M", "", "")
	convert100, tenths := convert(hundreds, 100, "C", "D", "M")
	convert10, units := convert(tenths, 10, "X", "L", "C")
	convert1, _ := convert(units, 1, "I", "V", "X")

	return convert1000 + convert100 + convert10 + convert1, nil
}
