package github_test

import (
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
