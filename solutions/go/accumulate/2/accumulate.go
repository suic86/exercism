package accumulate

func Accumulate(input []string, f func(string) string) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}
