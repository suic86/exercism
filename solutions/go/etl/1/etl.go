package etl

import "strings"

func Transform(scoreMap map[int][]string) map[string]int {
	result := make(map[string]int)
	for k, v := range scoreMap {
		for _, c := range v {
			result[strings.ToLower(c)] = k
		}
	}
	return result
}
