package cipher

import (
	"strings"
	"unicode"
)

//Cipher represent a cipher
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

//NewCaesar returns a Ceasar cipher
func NewCaesar() Cipher { return NewShift(3) }

type distance int

//NewShift returns a cipher with the given distance
func NewShift(n int) Cipher {
	if n != 0 && n < 26 && n > -26 {
		return distance(n)
	}
	return nil
}

func (d distance) Encode(s string) string {
	return strings.Map(func(r rune) rune {
		return mapping(r, int(d))
	}, s)
}

func (d distance) Decode(s string) string {
	return distance(-d).Encode(s)
}

func mapping(r rune, d int) rune {
	r = unicode.ToLower(r)
	if !unicode.IsLetter(r) {
		return -1
	}
	return (r-'a'+rune(26+d))%26 + 'a'
}

type vigenere string

//NewVigenere returns a cipher using the given key
func NewVigenere(key string) Cipher {
	validKey := true
	for _, c := range key {
		validKey = validKey && 'a' <= c && c <= 'z'
	}
	if validKey {
		return vigenere(key)
	}
	return nil
}

func (v vigenere) distance(index int) int {
	return int(rune(v[index%len(v)]) - 'a')
}

func (v vigenere) Encode(s string) (result string) {
	for i, r := range s {
		result += string(mapping(r, v.distance(i)))
	}
	return
}

func (v vigenere) Decode(s string) (result string) {
	for i, r := range s {
		result += string(mapping(r, -v.distance(i)))
	}
	return
}
