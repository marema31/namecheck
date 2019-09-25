package twitter_test

import (
	"testing"

	"github.com/marema31/namecheck/twitter"
)

func TestValidTwitterUsername(t *testing.T) {
	if !twitter.Check("marema31") || !twitter.Check("golang") {
		t.Error("Validation of valid username failed")
	}
}
func TestTwitterUsernameTooShort(t *testing.T) {
	if twitter.Check("") {
		t.Error("Validation of too short username failed")
	}
}
func TestTwitterUsernameTooLong(t *testing.T) {
	if twitter.Check("MyRidiculousLongName") {
		t.Error("Validation of too long username failed")
	}
}
func TestTwitterUsernameInvalidPattern(t *testing.T) {
	if twitter.Check("TheTwiTtEr") {
		t.Error("Validation of username containing pattern failed")
	}
}
func TestTwitterUsernameInvalidCharacters(t *testing.T) {
	if twitter.Check("是法国人") {
		t.Error("Validation of unicode username failed")
	}
}
