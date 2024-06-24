package goha

import (
	"bytes"
	"encoding/json"
	base_errors "errors"
	"github.com/genstackio/goha/errors"
	"github.com/lestrrat-go/jwx/jwt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var envs = map[string]string{
	"production": "https://api.helloasso.com",
	"sandbox":    "https://api.helloasso-sandbox.com",
}

func (hac *Client) Init(clientIdOrAccessToken string, clientKeyOrRefreshToken string, env string) {
	if len(clientIdOrAccessToken) > 450 || len(clientKeyOrRefreshToken) > 450 {
		hac.InitFromTokens(clientIdOrAccessToken, clientKeyOrRefreshToken, env)
	} else {
		hac.InitFromApiKey(clientIdOrAccessToken, clientKeyOrRefreshToken, env)
	}
}

func (hac *Client) InitFromApiKey(clientId string, clientSecret string, env string) {
	hac.InitCommon(env)
	hac.identity.Username = clientId
	hac.identity.Password = clientSecret
}

func (hac *Client) InitFromTokens(accessToken string, refreshToken string, env string) {
	hac.InitCommon(env)
	hac.tokens.AccessToken = accessToken
	hac.tokens.RefreshToken = refreshToken
}

func (hac *Client) InitCommon(env string) {
	hac.env = env
	endpoint, ok := envs[env]
	if !ok {
		endpoint = envs["production"]
	}
	hac.endpoint = endpoint
	hac.options = ClientOptions{
		MinExpirationDelay: 1,
	}
}

func (hac *Client) isIdentified() bool {
	return len(hac.identity.Username) > 0 || len(hac.tokens.AccessToken) > 0 || len(hac.tokens.RefreshToken) > 0
}

func (hac *Client) decodeToken(token string) (jwt.Token, error) {
	t, err := decodeJwt(token)

	return t, err
}
func (hac *Client) isTokenExpired(token string) bool {
	if len(token) <= 0 {
		return true
	}
	decodedToken, err := hac.decodeToken(token)
	if err != nil {
		return false
	}
	now := int64(math.Floor(float64(time.Now().UnixMilli()) / float64(1000)))
	exp := int64(math.Floor(float64(decodedToken.Expiration().UnixMilli()) / float64(1000)))
	return (exp - now) < (hac.options.MinExpirationDelay)
}
func (hac *Client) isAccessTokenValid() bool {
	if len(hac.tokens.AccessToken) <= 0 {
		return false
	}
	return !hac.isTokenExpired(hac.tokens.AccessToken)

}
func (hac *Client) createAuthTokensFromRefreshToken(clientId string, refreshToken string) (ClientTokens, error) {
	var ct ClientTokens
	if len(refreshToken) <= 0 {
		return ct, base_errors.New("empty refresh token, please reset tokens")
	}
	err := hac.postForm(
		"/oauth2/token",
		map[string]string{
			"client_id":     clientId,
			"refresh_token": refreshToken,
			"grant_type":    "refresh_token",
		},
		&ct,
	)
	return ct, err
}
func (hac *Client) createAuthTokensFromIdentity(identity ClientIdentity) (ClientTokens, error) {
	var ct ClientTokens
	if len(identity.Username) <= 0 {
		return ct, base_errors.New("empty identity, please set identity first")
	}
	err := hac.postForm(
		"/oauth2/token",
		map[string]string{
			"client_id":     identity.Username,
			"client_secret": identity.Password,
			"grant_type":    "client_credentials",
		},
		&ct,
	)
	return ct, err
}
func (hac *Client) refreshAccessToken() error {
	if hac.isTokenExpired(hac.tokens.RefreshToken) {
		tokens, err := hac.createAuthTokensFromIdentity(hac.identity)
		if err != nil {
			return err
		}
		hac.tokens = tokens
		return nil
	}
	tokens, err := hac.createAuthTokensFromRefreshToken(hac.identity.Username, hac.tokens.RefreshToken)
	if err != nil {
		return err
	}
	hac.tokens = tokens
	return nil
}
func (hac *Client) refreshAccessTokenIfNeeded() (bool, error) {
	if !hac.isTokenExpired(hac.tokens.AccessToken) {
		return false, nil
	}
	if hac.isTokenExpired(hac.tokens.RefreshToken) {
		tokens, err := hac.createAuthTokensFromIdentity(hac.identity)
		if err != nil {
			return false, err
		}
		hac.tokens = tokens
		return true, nil
	}
	tokens, err := hac.createAuthTokensFromRefreshToken(hac.identity.Username, hac.tokens.RefreshToken)
	if err != nil {
		return false, err
	}
	hac.tokens = tokens
	return true, nil
}
func (hac *Client) prepareAuthTokens() error {
	if !hac.isIdentified() {
		return nil
	}
	if hac.isAccessTokenValid() {
		return nil
	}
	return hac.refreshAccessToken()

}
func (hac *Client) request(uri string, method string, body interface{}, headers map[string]string, data interface{}) error {
	err := hac.prepareAuthTokens()
	if err != nil {
		return err
	}
	tokenHeaders := map[string]string{}
	if len(hac.tokens.AccessToken) > 0 {
		tokenHeaders["Authorization"] = "Bearer " + hac.tokens.AccessToken
	}
	finalHeaders := mergeMaps(headers, tokenHeaders)
	return hac.http(method, uri, body, finalHeaders, data)
}
func (hac *Client) http(method string, uri string, body interface{}, headers map[string]string, data interface{}) error {
	_, err := hac.fetch(
		hac.endpoint+uri,
		FetchOptions{
			Method:  method,
			Body:    body,
			Headers: headers,
			Options: HttpOptions{
				Timeout: 20,
			},
		},
		data,
	)
	return err
}
func (hac *Client) fetch(url string, opts FetchOptions, data interface{}) (*http.Response, error) {
	var bodyReader io.Reader = nil
	if nil != opts.Body {
		rawBody, err := json.Marshal(opts.Body)

		if err != nil {
			return nil, err
		}

		bodyReader = bytes.NewReader(rawBody)
	}

	req, err := http.NewRequest(opts.Method, url, bodyReader)

	if err != nil {
		return nil, err
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}
	client := http.Client{}
	if opts.Options.Timeout > 0 {
		client.Timeout = time.Duration(opts.Options.Timeout * 1000000000)
	}

	res, err := client.Do(req)

	err = extractErrorFromResponseIfNeeded(res, err, map[string]string{"url": url, "statusCode": strconv.Itoa(res.StatusCode)})

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(res.Body).Decode(data)

	return res, nil
}
func extractAccessDeniedError(err ErrorData, _ []byte, infos map[string]string) error {
	return errors.AccessDeniedError{Description: err.ErrorDescription, Url: infos["url"]}
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

	err2 := json.Unmarshal(respBytes, &errorData)

	if nil != err2 {
		return err2
	}

	switch errorData.Error {
	case "access_denied":
		return extractAccessDeniedError(errorData, respBytes, infos)
	default:
		return extractGenericError(errorData, respBytes, infos)
	}
}
func (hac *Client) postForm(uri string, vars map[string]string, data interface{}) error {
	vals := url.Values{}
	for k, v := range vars {
		vals.Set(k, v)
	}
	u := hac.endpoint + uri
	res, err := http.PostForm(u, vals)

	err = extractErrorFromResponseIfNeeded(res, err, map[string]string{"url": u, "statusCode": strconv.Itoa(res.StatusCode)})

	if err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(data)
}
func (hac *Client) createDocument(uri string, body interface{}, data interface{}) error {
	return hac.request(uri, http.MethodPost, body, map[string]string{"Content-Type": "application/json;charset=utf-8"}, data)
}
func (hac *Client) updateDocument(uri string, body interface{}, data interface{}) error {
	return hac.request(uri, http.MethodPut, body, map[string]string{"Content-Type": "application/json;charset=utf-8"}, data)
}
func (hac *Client) deleteDocument(uri string, body interface{}, data interface{}) error {
	return hac.request(uri, http.MethodDelete, body, map[string]string{"Content-Type": "application/json;charset=utf-8"}, data)
}
func (hac *Client) getDocument(uri string, data interface{}) error {
	return hac.request(uri, http.MethodGet, nil, map[string]string{"Content-Type": "application/json;charset=utf-8"}, data)
}
