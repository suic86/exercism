// Package grains is a solution for the Grains exercims task.
package grains

import "fmt"

// Square computes the number of grains on the given square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, fmt.Errorf("the square must be >= 1 and <= 64 but was %d", square)
	}
	return 1 << (square - 1), nil
}

// Total computes the total number of grains on the chessboard.
func Total() uint64 {
	return (1 << 64) - 1
}
