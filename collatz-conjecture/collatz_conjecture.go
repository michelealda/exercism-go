package collatzconjecture

import "errors"

func recursion(i, steps int) int {
	if i == 1 {
		return steps
	}
	if i%2 == 0 {
		return recursion(i/2, steps+1)
	}

	return recursion(3*i+1, steps+1)

}

// CollatzConjecture calculate the Collatz Conjecture of a positive number
func CollatzConjecture(i int) (int, error) {
	if i <= 0 {
		return 0, errors.New("Only positive numbers allowed")
	}

	return recursion(i, 0), nil
}
