package proverb

func ending(word string) string {
	return "And all for the want of a " + word + "."
}

func section(word1, word2 string) string {
	return "For want of a " + word1 + " the " + word2 + " was lost."
}

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	if len(rhyme) == 0 {
		return proverb
	}

	if len(rhyme) > 1 {
		for i, s := range rhyme[:len(rhyme)-1] {
			proverb = append(proverb, section(s, rhyme[i+1]))
		}
	}

	return append(proverb, ending(rhyme[0]))
}
