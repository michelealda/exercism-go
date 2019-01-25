package triangle

import "math"

//Kind represent the triangle type
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func isNotTriangle(a, b, c float64) bool {
	return math.IsNaN(a+b+c) ||
		math.IsInf(a+b+c, 0) ||
		a+b+c == 0 ||
		a+b < c ||
		b+c < a ||
		a+c < b
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if isNotTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c && c == a {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
