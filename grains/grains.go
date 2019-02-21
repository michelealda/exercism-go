package grains

import (
	"errors"
)

//Square calculates the number of grains
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("error")
	}
	return 1 << uint64(n-1), nil
}

//Total calculates the sum of the 64 squares
func Total() (sum uint64) {
	for i := 1; i <= 64; i++ {
		x, _ := Square(i)
		sum += x
	}
	return sum
}
