package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var alphanum, _ = regexp.Compile("^[a-zA-Z_0-9]*$")

func validate(username string) bool {
	if username == "" || utf8.RuneCountInString(username) > 15 {
		fmt.Println("Username must have 1 to 15 characters")
		return false
	}
	lusername := strings.ToLower(username)
	if strings.Contains(lusername, "twitter") {
		fmt.Println("Username must not contain \"Twitter\"")
		return false
	}
	if !alphanum.MatchString(username) {
		fmt.Println("Username must contains only alphanumerical characters")
		return false
	}
	return true
}

func main() {
	fmt.Println(validate("marema31"))
	fmt.Println(validate("golang"))
	fmt.Println(validate(""))
	fmt.Println(validate("MyRidiculousLongName"))
	fmt.Println(validate("TheTwiTtEr"))
	fmt.Println(validate("是法国人"))
}
