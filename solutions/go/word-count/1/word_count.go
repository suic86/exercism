package wordcount

import (
	"regexp"
	"strings"
)

// Frequency ...
type Frequency map[string]int

// WordCount ...
func WordCount(text string) Frequency {
	result := Frequency{}

	if len(text) == 0 {
		return result
	}

	re := regexp.MustCompile(`([a-zA-Z]+'?[a-zA-Z]+?|\d+)`)

	for _, word := range re.FindAllString(text, -1) {
		word = strings.ToLower(word)
		c, ok := result[word]
		if !ok {
			result[word] = 1
			continue
		}
		result[word] = c + 1
	}
	return result
}
