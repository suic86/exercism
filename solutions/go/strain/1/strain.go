package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(fn func(int) bool) Ints {
	var result Ints
	for _, v := range ints {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (ints Ints) Discard(fn func(int) bool) Ints {
	return ints.Keep(func(i int) bool { return !fn(i) })
}

func (lists Lists) Keep(fn func([]int) bool) Lists {
	var result Lists
	for _, v := range lists {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (strs Strings) Keep(fn func(string) bool) Strings {
	var result Strings
	for _, v := range strs {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
