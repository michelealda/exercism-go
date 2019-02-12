package lsproduct

import (
	"errors"
	"strconv"
)

func product(s string) (int, error) {
	p := 1
	for _, c := range s {
		x, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, err
		}
		p *= x
	}
	return p, nil
}

//LargestSeriesProduct returns the highest product of span consecutive digits
func LargestSeriesProduct(s string, span int) (int, error) {
	if span > len(s) || span < 0 {
		return 0, errors.New("invalid span")
	}
	if len(s) == 0 || span == 0 {
		return 1, nil
	}
	max := 0
	for i := range s[:len(s)-span+1] {
		p, err := product(s[i : i+span])
		if err != nil {
			return 0, err
		}
		if p > max {
			max = p
		}
	}
	return max, nil
}
