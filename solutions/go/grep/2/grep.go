package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	MultipleFiles bool
	// flags
	Inverted        bool
	LineNumbers     bool
	CaseInsensitive bool
	FileOnly        bool
	EntireLine      bool
}

func NewConfig(flags []string) Config {
	cfg := Config{}
	for _, f := range flags {
		switch f {
		case "-v":
			cfg.Inverted = true
		case "-n":
			cfg.LineNumbers = true
		case "-i":
			cfg.CaseInsensitive = true
		case "-x":
			cfg.EntireLine = true
		case "-l":
			cfg.FileOnly = true
		default:
			log.Fatalf("invalid flag: %s", f)
		}
	}
	return cfg
}

func SearchSingle(pattern string, config *Config, filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := []string{}
	sc := bufio.NewScanner(f)
	lnum := 0
	for sc.Scan() {
		lnum++
		row := sc.Text()
		match := false
		if config.EntireLine {
			if config.CaseInsensitive {
				match = strings.EqualFold(row, pattern)
			} else {
				match = row == pattern
			}
		} else {
			if config.CaseInsensitive {
				match = strings.Contains(strings.ToLower(row), strings.ToLower(pattern))
			} else {
				match = strings.Contains(row, pattern)
			}
		}
		if config.Inverted {
			match = !match
		}
		if match {
			if config.FileOnly {
				return []string{filename}
			}
			if config.LineNumbers {
				row = fmt.Sprintf("%d:%s", lnum, row)
			}
			if config.MultipleFiles {
				row = fmt.Sprintf("%s:%s", filename, row)
			}
			result = append(result, row)
		}
	}

	return result
}

func Search(pattern string, flags, files []string) []string {
	res := []string{}
	cfg := NewConfig(flags)
	cfg.MultipleFiles = len(files) > 1
	for _, f := range files {
		res = append(res, SearchSingle(pattern, &cfg, f)...)
	}
	return res
}
