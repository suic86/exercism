package diffsquares

// SumOfSquares compute the sum of integers from 0 to n.
// This is an O(1) implementation using a formula which is
// explained e.g. here: https://trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm#sumDifferences
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum computes the sum of integers from 0 to n
// This is an O(1) implementation using a formula which is
// explained e.g. here: https://www.mathsisfun.com/algebra/sequences-sums-arithmetic.html
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// Difference = SquareOfSum - SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
