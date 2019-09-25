package twitter_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/marema31/namecheck/twitter"
)

func TestValidTwitterUsername(t *testing.T) {
	var tw twitter.Twitter
	if !tw.Check("marema31") || !tw.Check("golang") {
		t.Error("Validation of valid username failed")
	}
}
func TestTwitterUsernameTooShort(t *testing.T) {
	var tw twitter.Twitter
	if tw.Check("") {
		t.Error("Validation of too short username failed")
	}
}
func TestTwitterUsernameTooLong(t *testing.T) {
	var tw twitter.Twitter
	if tw.Check("MyRidiculousLongName") {
		t.Error("Validation of too long username failed")
	}
}
func TestTwitterUsernameInvalidPattern(t *testing.T) {
	var tw twitter.Twitter
	if tw.Check("TheTwiTtEr") {
		t.Error("Validation of username containing pattern failed")
	}
}
func TestTwitterUsernameInvalidCharacters(t *testing.T) {
	var tw twitter.Twitter
	if tw.Check("是法国人") {
		t.Error("Validation of unicode username failed")
	}
}

type noHttpClient struct{}

func (n *noHttpClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(nil),
	}, nil
}

type yesHttpClient struct{}

func (y *yesHttpClient) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(nil),
	}, nil
}

func TestTwitterUsernameAvailable(t *testing.T) {
	var tw twitter.Twitter
	client := &noHttpClient{}

	if ok, _ := tw.IsAvailable(client, "toto"); !ok {
		t.Error("Validation of unicode username failed")
	}
}

func TestTwitterUsernameNotAvailable(t *testing.T) {
	var tw twitter.Twitter
	client := &yesHttpClient{}

	if ok, _ := tw.IsAvailable(client, "toto"); ok {
		t.Error("Validation of unicode username failed")
	}
}
