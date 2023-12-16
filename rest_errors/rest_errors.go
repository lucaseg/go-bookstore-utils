package rest_errors

import "net/http"

type RestError struct {
	Status  int
	Message string
	Error   string
}

func BadRequest(message string) *RestError {
	return &RestError{
		http.StatusBadRequest,
		message,
		"bad_request",
	}
}

func NotFound(message string) *RestError {
	return &RestError{
		http.StatusNotFound,
		message,
		"not_found",
	}
}

func InteralServerError(message string) *RestError {
	return &RestError{
		http.StatusInternalServerError,
		message,
		"internal_server_error",
	}
}
