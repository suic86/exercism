package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func SearchSingle(pattern string, flags []string, filename string, multi bool) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	inverted := slices.Contains(flags, "-v")
	linenum := slices.Contains(flags, "-n")
	caseins := slices.Contains(flags, "-i")
	fileonly := slices.Contains(flags, "-l")
	entireline := slices.Contains(flags, "-x")

	result := []string{}
	sc := bufio.NewScanner(f)
	lnum := 0
	for sc.Scan() {
		lnum++
		row := sc.Text()
		m := false
		if entireline {
			if caseins {
				m = strings.EqualFold(row, pattern)
			} else {
				m = row == pattern
			}
		} else {
			if caseins {
				m = strings.Contains(strings.ToLower(row), strings.ToLower(pattern))
			} else {
				m = strings.Contains(row, pattern)
			}
		}
		if inverted {
			m = !m
		}
		if m {
			if fileonly {
				return []string{filename}
			}
			if linenum {
				row = fmt.Sprintf("%d:%s", lnum, row)
			}
			if multi {
				row = fmt.Sprintf("%s:%s", filename, row)
			}
			result = append(result, row)
		}
	}

	return result
}

func Search(pattern string, flags, files []string) []string {
	res := []string{}
	for _, f := range files {
		res = append(res, SearchSingle(pattern, flags, f, len(files) > 1)...)
	}
	return res

}
