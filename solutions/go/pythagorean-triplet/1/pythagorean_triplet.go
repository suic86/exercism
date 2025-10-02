package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	r := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if isPythagoreanTriangle(a, b, c) {
					r = append(r, [3]int{a, b, c})
				}
			}
		}
	}
	return r
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	r := []Triplet{}
	for c := p; c > 0; c-- {
		for b := c - 1; b > 1; b-- {
			a := p - c - b
			if a > b || a > c {
				break
			}
			if isPythagoreanTriangle(a, b, c) {
				r = append(r, [3]int{a, b, c})
			}
		}
	}
	return r
}

func isPythagoreanTriangle(a, b, c int) bool {
	if a > b {
		b, a = a, b
	}
	if b > c {
		c, b = b, c
	}
	if a > b {
		b, a = a, b
	}
	return (a+b > c) && (a+c > b) && (b+c > a) && (a*a+b*b == c*c)
}
