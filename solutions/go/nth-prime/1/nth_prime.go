package prime

import (
	"errors"
	"math"
)

var cache = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

func getPrime(n int) int {
	if (n - 1) < len(cache) {
		return cache[n-1]
	}
	isPrime := func(n int) bool {
		if n%2 == 0 || n%3 == 0 || n%5 == 5 {
			return false
		}
		for i := 5; i < int(math.Sqrt(float64(n)))+1; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}
	for j := cache[len(cache)-1] + 2; ; j += 2 {
		if isPrime(j) {
			cache = append(cache, j)
			if (n - 1) < len(cache) {
				break
			}
		}
	}
	return cache[n-1]
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}
	return getPrime(n), nil
}
