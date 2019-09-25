package twitter

import (
	"net/http"
	"regexp"

	"github.com/marema31/namecheck/validate"
)

type Twitter struct{}

type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "Network Error"
}

func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}

var alphanum = regexp.MustCompile("^[a-zA-Z_0-9]*$")

const url = "https://twitter.com/"

func (t *Twitter) Check(username string) bool {
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

func (t *Twitter) IsAvailable(username string) (bool, error) {
	res, err := http.Get(url + username)
	if err != nil {
		return false, &ErrNetworkFailure{Cause: err}
	}
	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}
