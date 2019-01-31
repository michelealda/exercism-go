package strain

//Ints is a list of int
type Ints []int

//Lists is a list of lists
type Lists [][]int

//Strings is a list of strings
type Strings []string

//Keep the values matching the predicate
func (xs Ints) Keep(fn func(int) bool) (result Ints) {
	for _, x := range xs {
		if fn(x) {
			result = append(result, x)
		}
	}
	return
}

//Discard the values matching the predicate
func (xs Ints) Discard(fn func(int) bool) Ints {
	notFn := func(x int) bool {
		return !fn(x)
	}
	return xs.Keep(notFn)
}

//Keep the values matching the predicate
func (xs Lists) Keep(fn func([]int) bool) (result Lists) {
	for _, x := range xs {
		if fn(x) {
			result = append(result, x)
		}
	}
	return
}

//Keep the values matching the predicate
func (xs Strings) Keep(fn func(string) bool) (result Strings) {
	for _, x := range xs {
		if fn(x) {
			result = append(result, x)
		}
	}
	return
}
