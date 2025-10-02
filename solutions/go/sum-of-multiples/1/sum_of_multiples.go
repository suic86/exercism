package summultiples

func SumMultiples(limit int, divisors ...int) int {
	r := 0
	if len(divisors) == 0 {
		return r
	}
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d != 0 && i%d == 0 {
				r += i
				break
			}
		}
	}
	return r
}
