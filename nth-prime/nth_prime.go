package prime

func isDivisible(n int, prime []int) bool {
	for _, d := range prime {
		if n%d == 0 {
			return true
		}
	}
	return false
}

//Nth return the nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	prime := []int{2, 3, 5, 7, 11, 13, 17}
	if n < len(prime) {
		return prime[n-1], true
	}
	i := prime[len(prime)-1] + 1
	for {
		if isDivisible(i, prime) {
			i++
			continue
		}
		prime = append(prime, i)
		if n < len(prime) {
			return prime[n-1], true
		}
	}
}
