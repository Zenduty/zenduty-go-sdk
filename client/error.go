package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNoToken = errors.New("an empty token was provided")

	ErrAuthFailure = errors.New("failed to authenticate using the provided token")
)

type errorResponse struct {
	Error *Error `json:"error"`
}

type Error struct {
	ErrorResponse *Response
	Code          int         `json:"code,omitempty"`
	Errors        interface{} `json:"error,omitempty"`
	Message       string      `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s API call to %s failed %v. Errors: %s, Message: %s", e.ErrorResponse.Response.Request.Method, e.ErrorResponse.Response.Request.URL.String(), e.ErrorResponse.Response.Status, string(e.ErrorResponse.BodyBytes), e.Message)
}

// StatusCode returns the HTTP status code of the response that produced the
// error. The HTTP response is authoritative; Code is only a fallback because
// a JSON error body carrying its own "code" field overwrites it on decode.
func (e *Error) StatusCode() int {
	if e.ErrorResponse != nil && e.ErrorResponse.Response != nil {
		return e.ErrorResponse.Response.StatusCode
	}
	return e.Code
}

// StatusCode returns the HTTP status code carried by err when it is an API
// error from this client, and 0 otherwise (nil, transport failures, errors
// from other packages).
func StatusCode(err error) int {
	var apiErr *Error
	if errors.As(err, &apiErr) && apiErr != nil {
		return apiErr.StatusCode()
	}
	return 0
}

// IsNotFound reports whether err is an API error with HTTP status 404.
func IsNotFound(err error) bool {
	return StatusCode(err) == http.StatusNotFound
}
