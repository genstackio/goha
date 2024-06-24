package common

type MalformedError struct {
	Err error `json:"err,omitempty"`
}

func (err MalformedError) Error() string {
	return "Malformed request"
}

func (err MalformedError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       16993,
		StatusCode: 400,
	}
}
