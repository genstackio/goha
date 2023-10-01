package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func TestCredentialsCommand(args []string, env string) (int, error) {
	if len(args) < 2 {
		return 1, errors.New("<clientIdOrAccessToken> <clientSecretOrRefreshToken>")
	}
	clientId := args[0]
	clientSecret := args[1]
	ok, err := goha.TestApiCredentials(clientId, clientSecret, env)

	if !ok {
		return 3, errors.New("NOK - Bad credentials:" + err.Error())
	}

	fmt.Println("OK")

	return 0, nil
}
