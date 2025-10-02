// Package luhn implement the Luhn algorithm for credit card number validation.
package luhn

import (
	"unicode"
)

// Valid checks if a given card number if valid using Luhn algorithm
// see https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(cardNumber string) bool {

	// filter all digits
	digits := make([]int, 0, len(cardNumber))
	for _, v := range cardNumber {
		if unicode.IsDigit(v) {
			digits = append(digits, int(v)-'0')
		} else if !unicode.IsSpace(v) {
			return false
		}
	}

	if len(digits) <= 1 {
		return false
	}

	// get every second digit starting from right without reversing the slice
	// for card numbers with odd length every even index should be doubled
	// card: 12346
	//       ^ ^ ^
	// i:    01234
	// for card numbers with even length every even index should be doubled
	// card: 1234
	//        ^ ^
	// i:    0123
	sum, modulo := 0, 1-len(digits)%2
	for i, d := range digits {
		if i%2 == modulo {
			sum += d
			continue
		}

		d = 2 * d
		if d > 9 {
			d -= 9
		}

		sum += d
	}

	return sum%10 == 0
}
