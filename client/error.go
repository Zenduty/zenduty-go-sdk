package client

import (
	"errors"
	"fmt"
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
	return fmt.Sprintf("%s API call to %s failed %v. Code: %d, Errors: %v, Message: %s", e.ErrorResponse.Response.Request.Method, e.ErrorResponse.Response.Request.URL.String(), e.ErrorResponse.Response.Status, e.Code, e.Errors, e.Message)
}
