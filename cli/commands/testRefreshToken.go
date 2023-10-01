package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func TestRefreshTokenCommand(args []string, env string) (int, error) {
	if len(args) < 1 {
		return 1, errors.New("<refreshToken>")
	}
	refreshToken := args[0]
	ok, err := goha.TestApiRefreshToken(refreshToken, env)

	if !ok {
		return 3, errors.New("NOK - Bad credentials:" + err.Error())
	}

	fmt.Println("OK")

	return 0, nil
}
