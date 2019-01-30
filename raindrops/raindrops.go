package raindrops

import "strconv"

type factor struct {
	divisor int
	word    string
}

//Convert a
func Convert(n int) string {
	divisible := func(x, d int, s string) string {
		if x%d == 0 {
			return s
		}
		return ""
	}

	vocabulary := []factor{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	result := ""
	for _, x := range vocabulary {
		result += divisible(n, x.divisor, x.word)
	}

	if result != "" {
		return result
	}
	return strconv.Itoa(n)
}
