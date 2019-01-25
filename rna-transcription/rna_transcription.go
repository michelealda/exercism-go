package strand

import "strings"

//ToRNA calculates the RNA complement of a DNA strand
func ToRNA(dna string) string {
	mapping := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	return strings.Map(func(x rune) rune {
		return mapping[x]
	}, dna)
}
