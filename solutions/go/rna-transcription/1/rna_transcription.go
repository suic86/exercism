package strand

import "strings"

var conversion = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA convert DNA string to RNA string.
func ToRNA(dna string) string {
	var rna strings.Builder
	for _, n := range dna {
		rna.WriteRune(conversion[n])
	}
	return rna.String()
}
