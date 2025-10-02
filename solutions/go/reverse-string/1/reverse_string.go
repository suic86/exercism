package reverse

// Reverse reverses the provided string.
func Reverse(input string) string {
	runeArr := []rune(input)
	n := len(runeArr)
	for i := 0; i < n/2; i++ {
		runeArr[i], runeArr[n-1-i] = runeArr[n-1-i], runeArr[i]
	}
	return string(runeArr)
}
