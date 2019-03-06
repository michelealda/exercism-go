package allyourbase

import (
	"errors"
	"math"
)

//ConvertToBase convert a given number to the given base
func ConvertToBase(a int, digits []int, b int) ([]int, error) {
	if a < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if b < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	len := len(digits)
	if len == 0 {
		return []int{0}, nil
	}
	n := 0
	for i, d := range digits {
		if d < 0 || d >= a {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		n += d * int(math.Pow(float64(a), float64(len-1-i)))
	}
	if n < b {
		return []int{n}, nil
	}

	bs := []int{}
	for {
		q, r := n/b, n%b
		bs = append([]int{r}, bs...)
		n = q
		if n < b {
			bs = append([]int{n}, bs...)
			return bs, nil
		}
	}
}
