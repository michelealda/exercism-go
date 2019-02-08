package wordcount

import (
	"regexp"
	"strings"
)

//Frequency is a map for string occurence
type Frequency map[string]int

//WordCount return the frequency of each word in a sentence
func WordCount(input string) Frequency {
	re, _ := regexp.Compile(`[\w|']+`)
	words := re.FindAllString(strings.ToLower(input), -1)
	frequency := map[string]int{}
	for _, s := range words {
		frequency[strings.Trim(s, "'")]++
	}
	return frequency
}
