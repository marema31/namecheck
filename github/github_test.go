package github_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/marema31/namecheck/github"
)

func TestValidgithubUsername(t *testing.T) {
	var g github.Github
	if !g.Check("marema31") || !g.Check("golang") {
		t.Error("Validation of valid username failed")
	}
}
func TestgithubUsernameTooShort(t *testing.T) {
	var g github.Github
	if g.Check("") {
		t.Error("Validation of too short username failed")
	}
}
func TestgithubUsernameTooLong(t *testing.T) {
	var g github.Github
	if g.Check("My-RidiculousLongNameAndSomeIsTooLongOrToMuch") {
		t.Error("Validation of too long username failed")
	}
}
func TestgithubUsernameInvalidPattern(t *testing.T) {
	var g github.Github
	if g.Check("Thegithub") {
		t.Error("Validation of username containing pattern failed")
	}
}
func TestgithubUsernameInvalidCharacters(t *testing.T) {
	var g github.Github
	if g.Check("是法国人") {
		t.Error("Validation of unicode username failed")
	}
}

func TestgithubUsernameInvalidStartWith(t *testing.T) {
	var g github.Github
	if g.Check("-Thegithub") {
		t.Error("Validation of username containing pattern failed")
	}
}

func TestgithubUsernameInvalidEndWith(t *testing.T) {
	var g github.Github
	if g.Check("Thegithub-") {
		t.Error("Validation of username containing pattern failed")
	}
}

func TestgithubUsernameInvalidDoubleDash(t *testing.T) {
	var g github.Github
	if g.Check("The--github") {
		t.Error("Validation of username containing pattern failed")
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

func TestGithubUsernameAvailable(t *testing.T) {
	var tw github.Github
	client := &noHttpClient{}

	if ok, _ := tw.IsAvailable(client, "toto"); !ok {
		t.Error("Validation of unicode username failed")
	}
}

func TestGithubUsernameNotAvailable(t *testing.T) {
	var tw github.Github
	client := &yesHttpClient{}

	if ok, _ := tw.IsAvailable(client, "toto"); ok {
		t.Error("Validation of unicode username failed")
	}
}
