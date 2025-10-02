package hamming

import "fmt"

// Distance computes the hamming distance between two DNA segments a and b. The segments must have the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("The two DNA segments have different lenght (a: %d, b: %d)", len(a), len(b))
	}

	distance := 0

	for i, r := range a {
		if r != rune(b[i]) {
			distance++
		}
	}

	return distance, nil
}
