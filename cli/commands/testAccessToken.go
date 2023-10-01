package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func TestAccessTokenCommand(args []string, env string) (int, error) {
	if len(args) < 1 {
		return 1, errors.New("<accessToken>")
	}
	accessToken := args[0]
	ok, err := goha.TestApiAccessToken(accessToken, env)

	if !ok {
		return 3, errors.New("NOK - Bad credentials:" + err.Error())
	}

	fmt.Println("OK")

	return 0, nil
}
