package etl

import "strings"

//Transform the old map into the new score map
func Transform(input map[int][]string) map[string]int {
	result := map[string]int{}

	for score, values := range input {
		for _, v := range values {
			result[strings.ToLower(v)] = score
		}
	}

	return result
}
