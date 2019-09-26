package github

import (
	"net/http"
	"regexp"

	"github.com/marema31/namecheck/validate"
	"github.com/marema31/namecheck/webclient"
)

type Github struct{}

const url = "https://github.com/"

var alphanum = regexp.MustCompile("^[-a-zA-Z0-9]*$")
var beginWith = regexp.MustCompile("^[a-zA-Z0-9]")
var endWith = regexp.MustCompile("[a-zA-Z0-9]$")
var doubleDash = regexp.MustCompile("--$")

func (g *Github) Name() string {
	return "Github"
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

func (g *Github) IsAvailable(client webclient.Http, username string) (bool, error) {
	req, err := http.NewRequest("GET", url+username, nil)
	if err != nil {
		return false, err
	}

	// Use the web variable and not a http.Client to allow overriding
	res, err := client.Do(req)
	if err != nil {
		return false, &webclient.ErrNetworkFailure{Cause: err}
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}