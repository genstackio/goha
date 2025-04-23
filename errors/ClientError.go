package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type ClientError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err ClientError) Error() string {
	return "Client: " + err.Description + " (at " + err.Url + ")"
}

func (err ClientError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "malformed",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16998,
		StatusCode: 415,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
