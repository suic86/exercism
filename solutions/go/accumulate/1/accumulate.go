package accumulate

func Accumulate(input []string, f func(string) string) []string {
	result := make([]string, 0, len(input))
	for _, v := range input {
		result = append(result, f(v))
	}
	return result
}
