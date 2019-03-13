package armstrong

import (
	"math"
	"strconv"
)

//IsNumber returns true if it is an Armstrong number
func IsNumber(n int) bool {
	if n < 10 {
		return true
	}
	digits := strconv.Itoa(n)
	p := len(digits)
	for _, d := range digits {
		x, _ := strconv.Atoi(string(d))
		n -= int(math.Pow(float64(x), float64(p)))
	}
	return n == 0
}
