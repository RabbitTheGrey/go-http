package internal

import "net/http"

type Error struct {
	Message string
	Code    int
}

func (e *Error) Error() string {
	return e.Message
}

func ErrUnauthorized(msg string) *Error {
	return &Error{Message: msg, Code: http.StatusUnauthorized}
}

func ErrNotFound(msg string) *Error {
	return &Error{Message: msg, Code: http.StatusNotFound}
}

func ErrForbidden(msg string) *Error {
	return &Error{Message: msg, Code: http.StatusForbidden}
}

func ErrInternal(msg string) *Error {
	return &Error{Message: msg, Code: http.StatusInternalServerError}
}

func ErrBadRequest(msg string) *Error {
	return &Error{Message: msg, Code: http.StatusBadRequest}
}
