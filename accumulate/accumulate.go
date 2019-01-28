package accumulate

//Accumulate map each element using the given function
func Accumulate(xs []string, fn func(string) string) []string {
	result := []string{}
	for _, s := range xs {
		result = append(result, fn(s))
	}
	return result
}
