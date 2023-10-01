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
