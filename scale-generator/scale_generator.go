package scale

import "strings"

var sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var useFlatScale = map[string][]string{
	"F":  flatScale,
	"Bb": flatScale,
	"Eb": flatScale,
	"Ab": flatScale,
	"Db": flatScale,
	"Gb": flatScale,
	"d":  flatScale,
	"g":  flatScale,
	"c":  flatScale,
	"f":  flatScale,
	"bb": flatScale,
	"eb": flatScale,
}

// Scale generates the scale for the given tonic, interval
func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = strings.Repeat("m", 12)
	}

	scale, ok := useFlatScale[tonic]
	if !ok {
		scale = sharpScale
	}

	formattedTonic := normalizeTonic(tonic)
	start := findStart(formattedTonic, scale)

	result := []string{formattedTonic}
	for _, x := range interval[0 : len(interval)-1] {
		switch x {
		case 'M':
			start += 2
		case 'A':
			start += 3
		default:
			start++
		}
		result = append(result, scale[start%12])
	}

	return result
}

func findStart(item string, arr []string) int {
	for i := range arr {
		if arr[i] == item {
			return i
		}
	}
	return -1
}

func normalizeTonic(tonic string) string {
	if len(tonic) == 1 {
		return strings.ToUpper(tonic)
	}
	return strings.Title(tonic)
}
