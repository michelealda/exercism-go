package series

//All return all the sequence
func All(n int, s string) []string {
	r := []string{}
	for i := 0; i < len(s)-n+1; i++ {
		r = append(r, s[i:i+n])
	}
	return r
}

//UnsafeFirst returns the first sequence
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

//First returns the first sequence if ok
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
