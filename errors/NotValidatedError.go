package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type NotValidatedError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err NotValidatedError) Error() string {
	return "NotValidated: " + err.Description + " (at " + err.Url + ")"
}

func (err NotValidatedError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16999,
		StatusCode: 409,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
