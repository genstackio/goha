package commands

type command func(args []string, env string) (int, error)

var Commands = map[string]command{
	"list-organizations":                ListOrganizationsCommand,
	"list-organization-payments":        ListOrganizationPaymentsCommand,
	"test-access-token":                 TestAccessTokenCommand,
	"test-credentials":                  TestCredentialsCommand,
	"test-refresh-token":                TestRefreshTokenCommand,
	"test-tokens":                       TestTokensCommand,
	"test-tokens-and-organization-slug": TestTokensAndOrganizationSlugCommand,
}
