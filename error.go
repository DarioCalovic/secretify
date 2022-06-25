package secretify

import (
	"errors"
	"net/http"
)

var (
	// ErrGeneric is used for testing purposes and for errors handled later in the callstack
	ErrGeneric = errors.New("generic error")

	// ErrRecordNotFound (404) is returned when a record was not found
	ErrRecordNotFound = errors.New("record not found")

	ErrForbidden = errors.New("Forbidden to access resource")

	// ErrUnauthorized (401) is returned when user is not authorized
	ErrUnauthorized = errors.New(http.StatusText(http.StatusUnauthorized))
)
