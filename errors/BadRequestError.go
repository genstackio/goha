package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type BadRequestError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err BadRequestError) Error() string {
	return "BadRequest: " + err.Description + " (at " + err.Url + ")"
}

func (err BadRequestError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "malformed",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16997,
		StatusCode: 400,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
