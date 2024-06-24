package common

type JsonErrorResponse struct {
	Status     string                 `json:"status"`
	ErrorType  string                 `json:"errorType,omitempty"`
	Message    string                 `json:"message"`
	Detail     string                 `json:"detail,omitempty"`
	Code       int                    `json:"code"`
	StatusCode int                    `json:"statusCode"`
	Data       map[string]interface{} `json:"data,omitempty"`
}
