package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of runes in a string array
func ConcurrentFrequency(xs []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range xs {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	m := FreqMap{}
	for range xs {
		for r, n := range <-c {
			m[r] += n
		}
	}
	return m
}
