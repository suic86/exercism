package atbash

import "strings"

func Atbash(s string) string {
	var b strings.Builder
	g := 0
	for _, r := range s {
		switch {
		case 'a' <= r && r <= 'z':
			b.WriteRune('z' - r + 'a')
			g++
		case 'A' <= r && r <= 'Z':
			b.WriteRune('z' - r + 'A')
			g++
		case '0' <= r && r <= '9':
			b.WriteRune(r)
			g++
		default:
			continue
		}
		if g == 5 {
			g = 0
			b.WriteRune(' ')
		}
	}
	return strings.TrimRight(b.String(), " ")
}
