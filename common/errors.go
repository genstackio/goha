package common

import "strconv"

type AccessDeniedError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func (err AccessDeniedError) Error() string {
	return "Access denied: " + err.Description + "(at " + err.Url + ")"
}

func (err AccessDeniedError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
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

type GenericError struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	StatusCode  int    `json:"code,omitempty"`
}

func (err GenericError) Error() string {
	return "Unexpected error: " + err.Description + "(code: " + strconv.Itoa(err.StatusCode) + ", url: " + err.Url + ")"
}

func (err GenericError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
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
