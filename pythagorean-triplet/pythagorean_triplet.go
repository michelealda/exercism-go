package pythagorean

//Triplet represent a pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	ts := []Triplet{}
	for a := min; a <= max-2; a++ {
		for b := a + 1; b <= max-1; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return ts
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	ts := []Triplet{}
	for _, t := range Range(1, p/2) {
		if t[0]+t[1]+t[2] == p {
			ts = append(ts, t)
		}
	}
	return ts
}
