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
func Total() uint64 {
	//turns out the total of n squares is
	// (2**n)-1
	return (1 << 64) - 1
}
