package rest_errors

import (
	"fmt"
	"net/http"
)

type RestError interface {
	Status() int
	Message() string
	Error() string
	Causes() []interface{}
}

type restError struct {
	status  int           `json:"status"`
	message string        `json:"message"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restError) Status() int {
	return e.status
}
func (e restError) Message() string {
	return e.message
}

func (e restError) Causes() []interface{} {
	return e.causes
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.message, e.status, e.error, e.causes)
}

func NewRestError(status int, message string, error string, causes []interface{}) RestError {
	return restError{
		status:  status,
		message: message,
		error:   error,
		causes:  causes,
	}
}

func BadRequest(message string) RestError {
	return restError{
		status:  http.StatusBadRequest,
		message: message,
		error:   "bad_request",
	}
}

func NotFound(message string) RestError {
	return &restError{
		status:  http.StatusNotFound,
		message: message,
		error:   "not_found",
	}
}

func InternalServerError(message string) RestError {
	return &restError{
		status:  http.StatusInternalServerError,
		message: message,
		error:   "internal_server_error",
	}
}
