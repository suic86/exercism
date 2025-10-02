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

func SearchSingle(pattern string, config Config, filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	result := []string{}
	lnum := 0
	sc := bufio.NewScanner(f)

	if config.CaseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	for sc.Scan() {
		lnum++
		row := sc.Text()
		match := false
		line := row
		if config.CaseInsensitive {
			line = strings.ToLower(line)
		}
		if config.EntireLine {
			match = line == pattern
		} else {
			match = strings.Contains(line, pattern)
		}
		if match != config.Inverted {
			if config.FileOnly {
				result = append(result, filename)
				break
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
		res = append(res, SearchSingle(pattern, cfg, f)...)
	}
	return res
}
