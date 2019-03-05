package sieve

//Sieve calculates the sieve to a given limit
func Sieve(limit int) (result []int) {
	multiples := make([]bool, limit+1) //zero based
	for n := 2; n <= limit; {
		for x := n + n; x <= limit; x += n {
			multiples[x] = true
		}
		for n++; n <= limit && multiples[n]; n++ {

		}
	}

	for n := 2; n <= limit; n++ {
		if !multiples[n] {
			result = append(result, n)
		}
	}

	return
}
