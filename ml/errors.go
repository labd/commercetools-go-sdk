package ml

// Generated file, please do not change!!!

import (
	"fmt"
	"net/http"
)

type GenericRequestError struct {
	Content    []byte
	StatusCode int
	Response   *http.Response
}

func (e GenericRequestError) Error() string {
	return fmt.Sprintf("Request returned status code %d", e.StatusCode)
}
