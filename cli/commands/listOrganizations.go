package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func ListOrganizationsCommand(args []string, env string) (int, error) {
	if len(args) < 2 {
		return 1, errors.New("<clientId> <clientSecret>")
	}
	clientId := args[0]
	clientSecret := args[1]
	c := goha.Client{}
	c.Init(clientId, clientSecret, env)
	p, err := c.GetMyOrganizations()
	if err != nil {
		return 4, err
	}
	for _, v := range p.Items {
		fmt.Println(v.Name, v.City, v.OrganizationSlug, v.Url)
	}

	return 0, nil
}
