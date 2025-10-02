package armstrong

func IntPow(m, n int) int {
	r := m
	switch n {
	case 0:
		return 1
	case 1:
		return r
	default:
		for i := 2; i <= n; i++ {
			r *= m
		}
	}
	return r
}

func IsNumber(n int) bool {
	ds := []int{}
	o := n
	for o > 0 {
		ds = append(ds, (o % 10))
		o /= 10
	}
	l := len(ds)
	s := 0
	for _, d := range ds {
		s += IntPow(d, l)
	}
	return s == n
}
