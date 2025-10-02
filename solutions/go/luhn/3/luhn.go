// Package luhn implements the Luhn algorithm for credit card number validation.
package luhn

import (
	"unicode"
)

// Valid checks if a given card number if valid using Luhn algorithm
// see https://en.wikipedia.org/wiki/Luhn_algorithm
//
// This implementation using a single iteration over cardNumber.
// It uses the the following idea instead of reversing the string:
// To  get every second digit starting from right:
//
// 1. for card numbers with odd length every even index should be doubled
//     // card: 12346
//     //       ^ ^ ^
//     // i:    01234
// 2. for card numbers with even length every odd index should be doubled
//     // card: 1234
//     //        ^ ^
//     // i:    0123
//
// The function computes the sum for both odd and even length case
// and chooses the appropriate sum based on the length.
func Valid(cardNumber string) bool {

	var oddSum, evenSum, length, value, doubled int

	for _, v := range cardNumber {
		// spaces are allowed but must be stripped
		if unicode.IsSpace(v) {
			continue
		}

		// cardNumber must not contain non-digit character
		// other than space
		if !unicode.IsDigit(v) {
			return false
		}

		length++

		value = int(v) - '0'
		doubled = 2 * value

		if doubled >= 10 {
			doubled -= 9
		}

		if length%2 == 0 {
			evenSum += value
			oddSum += doubled
		} else {
			oddSum += value
			evenSum += doubled
		}

	}

	if length <= 1 {
		return false
	}

	if length%2 == 0 {
		return evenSum%10 == 0
	}

	return oddSum%10 == 0
}
