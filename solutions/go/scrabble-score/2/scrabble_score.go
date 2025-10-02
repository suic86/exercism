// Package scrabble provides solution for the Scrabble Score exercism.io task.
package scrabble

import "unicode"

// Score computes the Scrabble value for a given word.
func Score(word string) int {
	result := 0
	for _, v := range word {
		switch unicode.ToUpper(v) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}
	}
	return result
}
