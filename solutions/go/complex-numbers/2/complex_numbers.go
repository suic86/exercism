package complexnumbers

import "math"

type Number struct {
	real float64
	imag float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imag
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real: n1.real + n2.real,
		imag: n1.imag + n2.imag,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real: n1.real - n2.real,
		imag: n1.imag - n2.imag,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real: n1.real*n2.real - n1.imag*n2.imag,
		imag: n1.imag*n2.real + n1.real*n2.imag,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real: n.real * factor,
		imag: n.imag * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
		real: (n1.real*n2.real + n1.imag*n2.imag) / (n2.real*n2.real + n2.imag*n2.imag),
		imag: (n1.imag*n2.real - n1.real*n2.imag) / (n2.real*n2.real + n2.imag*n2.imag),
	}
}

func (n Number) Conjugate() Number {
	return Number{
		real: n.real,
		imag: -n.imag,
	}
}

func (n Number) Abs() float64 {
	return math.Hypot(n.real, n.imag)
}

func (n Number) Exp() Number {
	a := math.Exp(n.real)
	return Number{
		real: math.Cos(n.imag) * a,
		imag: math.Sin(n.imag) * a,
	}
}
