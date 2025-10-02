// Package raindrops provides an implementation of the Raindrops task in go track on exercism.io.
package raindrops

import "strconv"

// Convert computes the raindrops for a given number.
// The rules of raindrops are that if a given number:
//
// has 3 as a factor, add 'Pling' to the result.
// has 5 as a factor, add 'Plang' to the result.
// has 7 as a factor, add 'Plong' to the result.
// does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.  */
func Convert(number int) string {
	var result string
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(number)
	}
	return result
}
