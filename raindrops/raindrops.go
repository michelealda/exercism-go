package raindrops

import "fmt"

//Convert a
func Convert(n int) string {
	divisible := func(x, d int, s string) string {
		if x%d == 0 {
			return s
		}
		return ""
	}

	vocabulary := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	result := ""
	for k, v := range vocabulary {
		result = result + divisible(n, k, v)
	}

	if result != "" {
		return result
	}
	return fmt.Sprint(n)
}
