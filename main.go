package main

import (
	"fmt"

	"github.com/marema31/namecheck/twitter"
)

func main() {
	var t twitter.Twitter
	fmt.Println(t.Check("marema31"))
	fmt.Println(t.Check("golang"))
	fmt.Println(t.Check(""))
	fmt.Println(t.Check("MyRidiculousLongName"))
	fmt.Println(t.Check("TheTwcheckiTtEr"))
	fmt.Println(t.Check("是法国人"))
}
