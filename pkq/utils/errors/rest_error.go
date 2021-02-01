package projecterrors

import "net/http"

type RestError struct {
	Title  string
	Status int
	Detail string
}

func BadRequest(message string) *RestError {
	return &RestError{
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: message,
	}
}
