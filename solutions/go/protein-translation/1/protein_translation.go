package protein

import (
	"errors"
	"regexp"
)

// ErrInvalidBase is raised when invalid codon is detected
var ErrInvalidBase = errors.New("invalid codon")

// ErrStop is raised when stop codon in detected
var ErrStop = errors.New("stop codon")

var codonParser *regexp.Regexp = regexp.MustCompile("^[ACGU]{3}$")
var rnaParser *regexp.Regexp = regexp.MustCompile("[ACGU]{3}")

// FromCodon converts a codon to protein
func FromCodon(codon string) (string, error) {
	if !codonParser.MatchString(codon) {
		return "", ErrInvalidBase
	}

	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

}

// FromRNA converts an RNA string to an array of proteins
func FromRNA(rna string) ([]string, error) {
	var result []string
	for _, codon := range rnaParser.FindAllString(rna, -1) {
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return result, nil
			}
			return result, err
		}
		result = append(result, protein)
	}
	return result, nil
}
