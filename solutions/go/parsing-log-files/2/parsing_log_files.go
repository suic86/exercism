package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var (
	validator        = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`)
	splitter         = regexp.MustCompile(`<[~*=-]*>`)
	passwordMatcher  = regexp.MustCompile(`(?i)".*password.*"`)
	endOfLineMatcher = regexp.MustCompile(`end-of-line\d*`)
	userNameMatcher  = regexp.MustCompile(`User +(.+?) `)
)

func IsValidLine(text string) bool {
	return validator.MatchString(text)
}

func SplitLogLine(text string) []string {
	return splitter.Split(text, 3)
}

func CountQuotedPasswords(lines []string) int {
	res := 0
	for _, l := range lines {
		if passwordMatcher.MatchString(l) {
			res++
		}
	}
	return res
}

func RemoveEndOfLineText(text string) string {
	return endOfLineMatcher.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tl := []string{}
	for _, l := range lines {
		m := userNameMatcher.FindStringSubmatch(l)
		if m == nil {
			tl = append(tl, l)
			continue
		}
		tl = append(tl, fmt.Sprintf("[USR] %s %s", m[1], l))
	}
	return tl
}
