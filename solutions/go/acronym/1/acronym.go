package acronym

import (
	"regexp"
	"unicode"
)

func Abbreviate(s string) string {
	re := regexp.MustCompile(`[a-zA-Z']+`)
	words := re.FindAllString(s, -1)
	result := make([]rune, len(words))
	for i, w := range words {
		result[i] = unicode.ToUpper(rune(w[0]))
	}
	return string(result)
}
