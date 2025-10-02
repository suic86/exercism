package isogram

// IsIsogram checks if a given string is an isogram (i.e. a phrase without repeating letters).
// This implementation ignores alse non-alphabetic character and it's case insensitive.
// more information here: https://en.wikipedia.org/wiki/Isogram
func IsIsogram(text string) bool {
	var letters [26]bool
	var index int
	for _, r := range text {

		if r >= 'a' && r <= 'z' {
			index = int(r - 'a')
		} else if r >= 'A' && r <= 'Z' {
			index = int(r - 'A')
		} else {
			continue
		}

		if letters[index] {
			return false
		}

		letters[index] = true
	}
	return true
}
