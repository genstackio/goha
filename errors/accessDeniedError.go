package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type AccessDeniedError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err AccessDeniedError) Error() string {
	return "Access denied: " + err.Description + " (at " + err.Url + ")"
}

func (err AccessDeniedError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16993,
		StatusCode: 403,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
