package errors

import (
	goerror_errors "github.com/genstackio/goerror/errors"
)

type NotFoundError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err NotFoundError) Error() string {
	return "NotFound: " + err.Description + " (at " + err.Url + ")"
}

func (err NotFoundError) JsonResponse() goerror_errors.JsonErrorResponse {
	return goerror_errors.JsonErrorResponse{
		Status:     "error",
		ErrorType:  "not_found",
		Message:    err.Error(),
		Detail:     err.Description,
		Code:       16996,
		StatusCode: 404,
		Data: map[string]interface{}{
			"url": err.Url,
		},
	}
}
