package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	var numbers []int

	// fill the sieve
	for i := 0; i < limit+1; i++ {
		numbers = append(numbers, i)
	}

	result := []int{2}
	d := 2

	for {
		// mark multiples
		i := d + d
		for i <= limit {
			numbers[i] = -1
			i += d
		}

		// find next prime
		i = d + 1
		for i <= limit {
			if numbers[i] != -1 {
				d = i
				result = append(result, d)
				break
			}
			i += 1
		}

		if i >= limit {
			break
		}
	}

	return result
}
