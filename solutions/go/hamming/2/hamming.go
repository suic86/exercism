package hamming

import "fmt"

// Distance computes the hamming distance between two DNA segments a and b. The segments must have the same length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("the two DNA segments have different lenght (a: %d, b: %d)", len(a), len(b))
	}

	distance := 0
	br := []rune(b)

	for i, r := range a {
		if r != br[i] {
			distance++
		}
	}

	return distance, nil
}
