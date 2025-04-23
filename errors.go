package goha

import (
	"encoding/json"
	"github.com/genstackio/goha/errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

func extractAccessDeniedError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.AccessDeniedError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractForbiddenError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.ForbiddenError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractAuthenticationError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.AuthenticationError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractNotFoundError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.NotFoundError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractNotValidatedError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.NotValidatedError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractBadRequestError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.BadRequestError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractClientError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.ClientError{Description: err.ErrorDescription, Url: infos["url"]}
}
func extractGenericError(err ErrorData, _ []byte, infos map[string]string) error {
	statusCode, _ := strconv.Atoi(infos["statusCode"])
	return errors.GenericError{Description: err.ErrorDescription, Url: infos["url"], StatusCode: statusCode}
}

func extractErrorFromResponseIfNeeded(res *http.Response, err error, infos map[string]string) error {
	if err != nil {
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode < 400 {
		return nil
	}

	respBytes, _ := io.ReadAll(res.Body)

	errorData := ErrorData{}

	// we try to fetch response content as json, if failing, ignore the content
	_ = json.Unmarshal(respBytes, &errorData)

	switch errorData.Error {
	case "access_denied":
		return extractAccessDeniedError(errorData, respBytes, infos)
	default:
		switch res.StatusCode {
		case 400:
			return extractBadRequestError(errorData, respBytes, infos)
		case 401:
			return extractAuthenticationError(errorData, respBytes, infos)
		case 403:
			return extractForbiddenError(errorData, respBytes, infos)
		case 404:
			return extractNotFoundError(errorData, respBytes, infos)
		case 409:
			return extractNotValidatedError(errorData, respBytes, infos)
		case 415:
			return extractClientError(errorData, respBytes, infos)
		default:
			log.Println("HelloAsso API error response:", string(respBytes))
			return extractGenericError(errorData, respBytes, infos)
		}
	}
}
