package hamming

import "errors"

type runeTuple struct {
	first  rune
	second rune
}

func zip(a, b []rune) []runeTuple {
	result := make([]runeTuple, len(a))
	for i, x := range a {
		result[i] = runeTuple{x, b[i]}
	}
	return result
}

func filter(tuples []runeTuple, fn func(runeTuple) bool) []runeTuple {
	result := []runeTuple{}
	for _, t := range tuples {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}

func itemsAreDifferent(t runeTuple) bool {
	return t.first != t.second
}

// Distance calculates the Hamming Distance between two DNA sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0,
			errors.New("Strings are not the same lenght")
	}

	return len(filter(zip([]rune(a), []rune(b)),
			itemsAreDifferent)),
		nil
}
