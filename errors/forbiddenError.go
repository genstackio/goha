package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type ForbiddenError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err ForbiddenError) Error() string {
	return "Forbidden: " + err.Description + " (at " + err.Url + ")"
}

func (err ForbiddenError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16994,
		StatusCode: 403,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
