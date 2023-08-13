package goha

import (
	"github.com/lestrrat-go/jwx/jwt"
)

func decodeJwt(token string) (jwt.Token, error) {
	t, err := jwt.Parse([]byte(token))

	return t, err
}

func mergeMaps(m1 map[string]string, m2 map[string]string) map[string]string {
	merged := make(map[string]string)
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

func TestApiCredentials(clientId string, clientSecret string, env string) (bool, error) {
	c := Client{}
	c.Init(clientId, clientSecret, env)
	_, err := c.GetMyOrganizations()

	if err != nil {
		return false, err
	}

	return true, nil
}
