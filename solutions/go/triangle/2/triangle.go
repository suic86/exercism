package triangle

import "math"

type Kind string

const (
	NaT = Kind("not a triangle")
	Equ = Kind("equilateral")
	Iso = Kind("isosceles")
	Sca = Kind("scalene")
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	switch {
	// check triangle inequality
	case isInvalidValue(a) || isInvalidValue(b) || isInvalidValue(c):
		return NaT
	case (a+b) < c || (a+c) < b || (b+c) < a:
		return NaT
	case (a == b) && (b == c):
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}

func isInvalidValue(v float64) bool {
	return v <= 0 || math.IsNaN(v) || math.IsInf(v, 1) || math.IsInf(v, -1)
}
