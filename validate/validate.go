package validate

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

func LengthLimit(username string, min, max int) bool {
	l := utf8.RuneCountInString(username)
	if l < min || l > max {
		fmt.Printf("Username must have %d to %d characters\n", min, max)
		return false
	}
	return true
}

func IllegalPatterns(username string, patterns []string) bool {
	lusername := strings.ToLower(username)
	for _, pattern := range patterns {
		if strings.Contains(lusername, pattern) {
			fmt.Printf("Username must not contain \"%s\"\n", pattern)
			return false
		}
	}
	return true
}

func IllegalChars(username string, pattern *regexp.Regexp) bool {
	if !pattern.MatchString(username) {
		fmt.Printf("Username must contain only \"%s\"\n", pattern)
		return false
	}
	return true
}
