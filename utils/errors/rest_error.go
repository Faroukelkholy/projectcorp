package projectErrors


import "net/http"

type RestErr struct {
	Title string
	Status  int
	Detail   string
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Title:   "Bad Request",
		Status:  http.StatusBadRequest,
		Detail: message,

	}
}


