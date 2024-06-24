package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
	"strconv"
)

type GenericError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	StatusCode  int    `json:"code,omitempty"`
}

func (err GenericError) Error() string {
	return "Unexpected error: " + err.Description + "(code: " + strconv.Itoa(err.StatusCode) + ", url: " + err.Url + ")"
}

func (err GenericError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "unexpected",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16992,
		StatusCode: err.StatusCode,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
