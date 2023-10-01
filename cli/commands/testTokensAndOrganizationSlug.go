package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func TestTokensAndOrganizationSlugCommand(args []string, env string) (int, error) {
	if len(args) < 3 {
		return 1, errors.New("<accessToken> <refreshToken> <organizationSlug>")
	}
	accessToken := args[0]
	refreshToken := args[1]
	organizationSlug := args[2]
	ok, err := goha.TestApiAccessAndRefreshTokensAndOrganizationSlug(accessToken, refreshToken, organizationSlug, env)

	if !ok {
		return 3, errors.New("NOK - Bad credentials:" + err.Error())
	}

	fmt.Println("OK")

	return 0, nil
}
