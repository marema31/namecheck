package github

import (
	"net/http"
	"regexp"

	"github.com/marema31/namecheck/validate"
)

type Github struct{}

const url = "https://github.com/"

var alphanum = regexp.MustCompile("^[-a-zA-Z0-9]*$")
var beginWith = regexp.MustCompile("^[a-zA-Z0-9]")
var endWith = regexp.MustCompile("[a-zA-Z0-9]$")
var doubleDash = regexp.MustCompile("--$")

type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "Network Error"
}

func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}

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

func (g *Github) IsAvailable(username string) (bool, error) {
	res, err := http.Get(url + username)
	if err != nil {
		return false, &ErrNetworkFailure{Cause: err}
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}
