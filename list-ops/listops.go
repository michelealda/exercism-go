package listops

// IntList represent a list of integers
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldr aggregates a list folding from the head to the tail
func (list IntList) Foldr(fn binFunc, acc int) int {
	for _, e := range list {
		acc = fn(acc, e)
	}
	return acc
}

// Foldl aggregates a list folding from the tail to the head
func (list IntList) Foldl(fn binFunc, acc int) int {
	return list.Reverse().Foldr(fn, acc)
}

// Length return the lenght of the list
func (list IntList) Length() int {
	return len(list)
}

// Filter the list with the given predicate
func (list IntList) Filter(fn predFunc) IntList {
	r := []int{}
	for _, e := range list {
		if fn(e) {
			r = append(r, e)
		}
	}
	return IntList(r)
}

// Map returns a list where the function is applied to each element
func (list IntList) Map(fn unaryFunc) IntList {
	r := []int{}
	for _, e := range list {
		r = append(r, fn(e))
	}
	return IntList(r)
}

// Reverse the list
func (list IntList) Reverse() IntList {
	r := []int{}
	for _, e := range list {
		r = append([]int{e}, r...)
	}
	return IntList(r)
}

// Append to the end of the list
func (list IntList) Append(items IntList) IntList {
	return append(list, items...)
}

// Concat the given list to the end
func (list IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		list = list.Append(l)
	}
	return list
}
