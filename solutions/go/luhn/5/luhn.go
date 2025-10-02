// Package luhn implements the Luhn algorithm for credit card number validation.
package luhn

import (
	"strings"
	"unicode"
)

// Valid checks if a given card number if valid using Luhn algorithm.
// See https://en.wikipedia.org/wiki/Luhn_algorithm for more information.
func Valid(cardNumber string) bool {

	cardNumber = strings.ReplaceAll(cardNumber, " ", "")

	if len(cardNumber) <= 1 {
		return false
	}

	var value int

	isDoubled := len(cardNumber)%2 == 0
	sum := 0

	for _, v := range cardNumber {
		if !unicode.IsDigit(v) {
			return false
		}

		value = int(v) - '0'

		if isDoubled {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
		isDoubled = !isDoubled
	}

	return sum%10 == 0
}
