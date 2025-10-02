package anagram

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	s := []byte(strings.ToLower(str))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return string(s)
}

func Detect(word string, candidates []string) []string {
	normalized := strings.ToLower(word)
	sorted := sortString(word)
	length := len(word)

	result := make([]string, 0, len(candidates))
	for _, w := range candidates {
		if len(w) != length || strings.ToLower(w) == normalized {
			continue
		}
		if sortString(w) == sorted {
			result = append(result, w)
		}
	}
	return result
}
