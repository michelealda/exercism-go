package anagram

import (
	"sort"
	"strings"
)

func stringToInts(s string) []int {
	ints := make([]int, len(s))

	for i, c := range s {
		ints[i] = int(c - 0)
	}
	sort.Ints(ints)
	return ints
}

func isAnagram(s, c string) bool {
	if len(s) != len(c) || s == c {
		return false
	}
	rs, rc := stringToInts(s), stringToInts(c)

	for i := range rs {
		if rs[i] != rc[i] {
			return false
		}
	}
	return true
}

//Detect anagrams of the given word
func Detect(subject string, candidates []string) []string {
	results := []string{}

	for _, c := range candidates {
		if isAnagram(strings.ToLower(subject), strings.ToLower(c)) {
			results = append(results, c)
		}
	}
	return results
}
