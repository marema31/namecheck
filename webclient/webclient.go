package webclient

import (
	"net/http"
)

type Http interface {
	Do(*http.Request) (*http.Response, error)
}

type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "Network Error"
}

func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}
