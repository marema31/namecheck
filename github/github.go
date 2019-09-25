package github

import (
	"regexp"

	"github.com/marema31/namecheck/validate"
)

type Github struct{}

var alphanum = regexp.MustCompile("^[-a-zA-Z0-9]*$")
var beginWith = regexp.MustCompile("^[a-zA-Z0-9]")
var endWith = regexp.MustCompile("[a-zA-Z0-9]$")
var doubleDash = regexp.MustCompile("--$")

func (g *Github) Check(username string) bool {
	if !validate.LengthLimit(username, 1, 39) {
		return false
	}
	if !validate.IllegalChars(username, alphanum) {
		return false
	}
	if !validate.IllegalChars(username, beginWith) {
		return false
	}
	if !validate.IllegalChars(username, endWith) {
		return false
	}
	if validate.IllegalChars(username, doubleDash) {
		return false
	}
	return true
}
