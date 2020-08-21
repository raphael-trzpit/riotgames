package api

import "net/http"

type Error struct {
	msg      string
	response *http.Response
}

func NewError(response *http.Response) error {
	if response == nil {
		return nil
	}

	if response.StatusCode >= 400 {
		return &Error{
			msg:      response.Status,
			response: response,
		}
	}

	return nil
}

func (e *Error) Response() *http.Response {
	return e.response
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) As(otherE error) bool {
	otherApiE, ok := otherE.(*Error)
	if !ok {
		return false
	}

	if e.response == nil && otherApiE.response == nil {
		return true
	}

	if e.response == nil || otherApiE.response == nil {
		return false
	}

	return otherApiE.response.Status == e.response.Status
}
