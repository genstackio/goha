package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func TestTokensCommand(args []string, env string) (int, error) {
	if len(args) < 2 {
		return 1, errors.New("<accessToken> <refreshToken>")
	}
	accessToken := args[0]
	refreshToken := args[1]
	ok, err := goha.TestApiAccessAndRefreshTokens(accessToken, refreshToken, env)

	if !ok {
		return 3, errors.New("NOK - Bad credentials:" + err.Error())
	}

	fmt.Println("OK")

	return 0, nil
}
