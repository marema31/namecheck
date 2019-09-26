package checker

import "github.com/marema31/namecheck/webclient"

type Checker interface {
	Check(username string) bool
	Name() string
	IsAvailable(client webclient.Http, username string) (bool, error)
}
