package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type frequency struct {
	char   rune
	length int
}

// RunLengthEncode ...
func RunLengthEncode(input string) string {
	if len(input) < 2 {
		return input
	}

	char := rune(input[0])
	count := 1

	var result strings.Builder

	for _, r := range input[1:] {
		if r == char {
			count++
			continue
		}

		if count == 1 {
			result.WriteRune(char)
		} else {
			result.WriteString(fmt.Sprintf("%d%c", count, char))
		}
		char = r
		count = 1
	}

	if count == 1 {
		result.WriteRune(char)
	} else {
		result.WriteString(fmt.Sprintf("%d%c", count, char))
	}

	return result.String()
}

var parser = regexp.MustCompile(`(\d*?)([a-zA-Z ]+?)`)

func RunLengthDecode(input string) string {
	var length int
	var char string
	var result strings.Builder

	for _, m := range parser.FindAllStringSubmatch(input, -1) {
		length = 1
		if m[1] != "" {
			length, _ = strconv.Atoi(m[1])
		}

		char = m[2]
		result.WriteString(strings.Repeat(char, length))
	}
	return result.String()
}
