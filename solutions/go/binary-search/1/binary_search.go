package binarysearch

// SearchInts implements binary search algorithm.
func SearchInts(list []int, key int) int {
	min := 0
	max := len(list) - 1

	for max >= min {
		i := (min + max) / 2
		guess := list[i]
		if guess == key {
			return i
		}
		if guess < key {
			min = i + 1
			continue
		}
		if guess > key {
			max = i - 1
			continue
		}
	}

	return -1
}
