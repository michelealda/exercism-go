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

func concurrentFreq(xs []string, c chan FreqMap) {
	for _, s := range xs {
		c <- Frequency(s)
	}
	close(c)
}

// ConcurrentFrequency counts the frequency of runes in a string array
func ConcurrentFrequency(xs []string) FreqMap {
	c := make(chan FreqMap)
	go concurrentFreq(xs, c)
	m := FreqMap{}
	for v := range c {
		for r, n := range v {
			m[r] += n
		}
	}
	return m
}
