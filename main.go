package main

import (
	"fmt"
	"log"

	"github.com/marema31/namecheck/github"
	"github.com/marema31/namecheck/twitter"
)

func main() {
	var t twitter.Twitter
	var g github.Github

	usernames := []string{"golang", "", "MyRidiculousLongName", "TheTwitTeR", "是法国人", "marema31", "marema311", "marema3111"}

	for _, username := range usernames {
		fmt.Printf("User %s is \n", username)

		fmt.Printf("    Twitter: ")
		if t.Check(username) {
			fmt.Printf("valid")
			available, err := t.IsAvailable(username)
			if err != nil {
				log.Printf("No way to contact Twitter: %s", err)

				// On peut aussi redeclarer:
				// type wrapper interface {
				//    Unwrap() error
				//}
				// et remplacer le if par err, ok := err.(wrapper); ok
				if err, ok := err.(interface{ Unwrap() error }); ok {
					log.Fatal(err.Unwrap())
				}
			}
			if available {
				fmt.Println(", available")
			} else {
				fmt.Println(", not available")
			}
		} else {
			fmt.Println("invalid")
		}

		fmt.Printf("    Github:  ")
		if g.Check(username) {
			fmt.Printf("valid")
			available, err := t.IsAvailable(username)
			if err != nil {
				log.Fatalf("No way to contact Twitter: %V", err)
			}
			if available {
				fmt.Println(", available")
			} else {
				fmt.Println(", not available")
			}

		} else {
			fmt.Println("invalid")
		}
	}
}
