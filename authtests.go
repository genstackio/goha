package goha

func TestApiCredentials(clientIdOrAccessToken string, clientSecretOrRefreshToken string, env string) (bool, error) {
	c := Client{}
	c.Init(clientIdOrAccessToken, clientSecretOrRefreshToken, env)
	_, err := c.GetMyOrganizations()

	if err != nil {
		return false, err
	}

	return true, nil
}

func TestApiAccessToken(accessToken string, env string) (bool, error) {
	c := Client{}
	c.InitFromTokens(accessToken, "", env)
	_, err := c.GetMyOrganizations()

	if err != nil {
		return false, err
	}

	return true, nil
}

func TestApiRefreshToken(refreshToken string, env string) (bool, error) {
	c := Client{}
	c.InitFromTokens("", refreshToken, env)
	_, err := c.GetMyOrganizations()

	if err != nil {
		return false, err
	}

	return true, nil
}

func TestApiAccessAndRefreshTokens(accessToken string, refreshToken string, env string) (bool, error) {
	c := Client{}
	c.InitFromTokens(accessToken, refreshToken, env)
	_, err := c.GetMyOrganizations()

	if err != nil {
		return false, err
	}

	return true, nil
}

func TestApiAccessAndRefreshTokensAndOrganizationSlug(accessToken string, refreshToken string, organizationSlug string, env string) (bool, error) {
	c := Client{}
	c.InitFromTokens(accessToken, refreshToken, env)
	_, err := c.GetOrganizationPayments(organizationSlug, GetPaymentsOptions{})

	if err != nil {
		return false, err
	}

	return true, nil
}
