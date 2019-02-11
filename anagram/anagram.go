package anagram

import (
	"sort"
	"strings"
)

func orderStringCharacters(s string) []string {
	c := strings.Split(s, "")
	sort.Strings(c)
	return c
}

func isAnagram(s, c string) bool {
	if len(s) != len(c) || s == c {
		return false
	}
	rs, rc := orderStringCharacters(s), orderStringCharacters(c)

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
