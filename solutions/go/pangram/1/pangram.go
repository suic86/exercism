package pangram

func IsPangram(text string) bool {
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
		letters[index] = true
	}
	for _, v := range letters {
		if !v {
			return false
		}
	}
	return true
}
