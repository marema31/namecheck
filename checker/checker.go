package checker

import "github.com/marema31/namecheck/webclient"

type Checker interface {
	Check(username string) bool
	IsAvailable(client webclient.Http, username string) (bool, error)
}
