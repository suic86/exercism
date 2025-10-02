// Package listops implement some list operation function.
// It provides a solution for exercism.io's List Ops task
package listops

// IntList represent a list of integers.
type IntList []int

// Append joins two IntLists.
func (list IntList) Append(other IntList) IntList {
	return append(list, other...)
}

// Concat appends multiple IntList to an IntList
func (list IntList) Concat(others []IntList) IntList {
	result := list
	for _, v := range others {
		result = result.Append(v)
	}
	return result
}

// Length return the length of the IntLinst
func (list IntList) Length() int {
	return len(list)
}

// Reverse reverses the IntList
func (list IntList) Reverse() IntList {
	result := make([]int, len(list), len(list))
	length := len(list)
	for i, v := range list {
		result[length-i-1] = v
	}
	return result
}

// Foldl implements the fold left higher order function
func (list IntList) Foldl(fn func(int, int) int, init int) int {
	aggregator := init
	for _, v := range list {
		aggregator = fn(aggregator, v)
	}
	return aggregator
}

// Foldr implements the fold right higher order function.
func (list IntList) Foldr(fn func(int, int) int, init int) int {
	aggregator := init
	for _, v := range list.Reverse() {
		aggregator = fn(v, aggregator)
	}
	return aggregator
}

// Map applies the function fn to every element of the list.
func (list IntList) Map(fn func(int) int) IntList {
	result := make([]int, len(list), len(list))
	for i, v := range list {
		result[i] = fn(v)
	}
	return IntList(result)
}

// Filter return a list of element from the list which satify the given predicate.
func (list IntList) Filter(pred func(int) bool) IntList {
	result := make([]int, 0, len(list))
	for _, v := range list {
		if pred(v) {
			result = append(result, v)
		}
	}
	return IntList(result)
}
