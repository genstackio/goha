package goha

type Client struct {
	endpoint string
	env      string
	identity ClientIdentity
	tokens   ClientTokens
	options  ClientOptions
}

type ClientOptions struct {
	MinExpirationDelay int64
}
type ClientIdentity struct {
	Username string
	Password string
}
type ClientTokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int64  `json:"expires_in,omitempty"`
}

type FetchOptions struct {
	Method  string
	Body    interface{}
	Headers map[string]string
	Options HttpOptions
}

type Token struct {
	Exp int64
}

type HttpOptions struct {
	Timeout int64
}

type ErrorData struct {
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
