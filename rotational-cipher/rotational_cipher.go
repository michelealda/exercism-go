package rotationalcipher

import "unicode"

const plain = "abcdefghijklmnopqrstuvwxyz"

//RotationalCipher encode a string by rotating the alphabet
func RotationalCipher(s string, x int) (result string) {
	if x == 0 {
		return s
	}

	getIndex := func(r, offset rune, x int) int {
		return (int(r-offset) + x) % 26
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			result += string(r)
		}
		index := 0
		if r >= 'a' && r <= 'z' {
			index = getIndex(r, 'a', x)
			result += string(plain[index])
		}
		if r >= 'A' && r <= 'Z' {
			index = getIndex(r, 'A', x)
			result += string(unicode.ToUpper(rune(plain[index])))
		}
	}
	return
}
