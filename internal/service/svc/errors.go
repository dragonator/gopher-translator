package svc

import "net/http"

// Errors -
var (
	ErrDecodeRequest = &Error{StatusCode: http.StatusBadRequest, Message: "request decoding failed"}
	ErrInvalidInput  = &Error{StatusCode: http.StatusBadRequest, Message: "invalid input"}
)

// Error -
type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	return e.Message
}
