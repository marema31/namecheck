package validate_test

import (
	"regexp"
	"testing"

	"github.com/marema31/namecheck/validate"
)

func TestValidLengthUsername(t *testing.T) {
	if !validate.LengthLimit("toto12", 5, 8) {
		t.Error("Validation of valid username failed")
	}
}
func TestLengthUsernameTooShort(t *testing.T) {
	if validate.LengthLimit("toto", 5, 8) {
		t.Error("Validation of too short username failed")
	}
}
func TestLengthUsernameTooLong(t *testing.T) {
	if validate.LengthLimit("toto12", 2, 4) {
		t.Error("Validation of too long username failed")
	}
}

func TestUsernameValidPattern(t *testing.T) {
	if !validate.IllegalPatterns("TheTwiTtEr", []string{"not"}) {
		t.Error("Validation of username containing valid pattern failed")
	}
}

func TestUsernameInvalidPattern(t *testing.T) {
	if validate.IllegalPatterns("TheNotTwiTtEr", []string{"not"}) {
		t.Error("Validation of username containing invalid pattern failed")
	}
}

func TestUsernameValidCharacters(t *testing.T) {
	var alphanum = regexp.MustCompile("^[a-zA-Z_0-9]*$")
	if !validate.IllegalChars("TheLatestAndBiggestInformations", alphanum) {
		t.Error("Validation of unicode username failed")
	}
}

func TestUsernameInvalidCharacters(t *testing.T) {
	var alphanum = regexp.MustCompile("^[a-zA-Z_0-9]*$")
	if validate.IllegalChars("是法国人", alphanum) {
		t.Error("Validation of unicode username failed")
	}
}

func BenchmarkIllegalChars(b *testing.B) {
	var alphanum = regexp.MustCompile("^[a-zA-Z_0-9]*$")
	for n := 0; n < b.N; n++ {
		_ = validate.IllegalChars("TheLatestAndBiggestInformations", alphanum)
	}
}
