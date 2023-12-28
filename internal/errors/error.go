package errors

import "net/http"

type Error struct {
	code    int
	message string
}

func New(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) HTTPStatusCode() int {
	switch {
	case e.code >= CodeAuth:
		return http.StatusUnauthorized
	}
	return http.StatusBadRequest
}
