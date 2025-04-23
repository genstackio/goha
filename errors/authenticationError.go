package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type AuthenticationError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err AuthenticationError) Error() string {
	return "Authentication: " + err.Description + " (at " + err.Url + ")"
}

func (err AuthenticationError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16995,
		StatusCode: 401,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
