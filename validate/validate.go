package validate

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func LengthLimit(username string, min, max int) bool {
	l := utf8.RuneCountInString(username)
	if l < min || l > max {
		return false
	}
	return true
}

func IllegalPatterns(username string, patterns []string) bool {
	lusername := strings.ToLower(username)
	for _, pattern := range patterns {
		if strings.Contains(lusername, pattern) {
			return false
		}
	}
	return true
}

func IllegalChars(username string, pattern *regexp.Regexp) bool {
	if !pattern.MatchString(username) {
		return false
	}
	return true
}
