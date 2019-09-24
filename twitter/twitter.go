package twitter

import (
	"regexp"

	"github.com/marema31/namecheck/validate"
)

var alphanum = regexp.MustCompile("^[a-zA-Z_0-9]*$")

func Check(username string) bool {
	if !validate.LengthLimit(username, 1, 15) {
		return false
	}
	if !validate.IllegalPatterns(username, []string{"twitter"}) {
		return false
	}
	if !validate.IllegalChars(username, alphanum) {
		return false
	}
	return true
}
