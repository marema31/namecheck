package twitter

import (
	"net/http"
	"regexp"

	"github.com/marema31/namecheck/validate"
	"github.com/marema31/namecheck/webclient"
)

type Twitter struct{}

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

func (t *Twitter) IsAvailable(client webclient.Http, username string) (bool, error) {
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
