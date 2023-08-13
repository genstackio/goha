package goha

type Client struct {
	clientId     string
	clientSecret string
	endpoint     string
	env          string
	identity     ClientIdentity
	tokens       ClientTokens
	options      ClientOptions
}

type ClientOptions struct {
	MinExpirationDelay int64
}
type ClientIdentity struct {
	Username string
	Password string
}
type ClientTokens struct {
	AccessToken  string
	RefreshToken string
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
