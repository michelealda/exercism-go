package summultiples

func isMultipleOfAny(x int, divisors ...int) bool {
	for _, d := range divisors {
		if d != 0 && x%d == 0 {
			return true
		}
	}
	return false
}

//SumMultiples sums all multiple of given divisors
func SumMultiples(limit int, divisors ...int) (sum int) {
	for n := 1; n < limit; n++ {
		if isMultipleOfAny(n, divisors...) {
			sum += n
		}
	}
	return
}
