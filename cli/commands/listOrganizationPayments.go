package commands

import (
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func ListOrganizationPaymentsCommand(args []string, env string) (int, error) {
	if len(args) < 3 {
		return 1, errors.New("<clientId> <clientSecret> <organizationSlug>")
	}
	clientId := args[0]
	clientSecret := args[1]
	c := goha.Client{}
	c.Init(clientId, clientSecret, env)
	p, err := c.GetOrganizationPayments(args[2], goha.GetPaymentsOptions{})
	if err != nil {
		return 5, err
	}
	for _, v := range p.Items {
		fmt.Println(v.Id, v.Amount, v.Date)
	}

	return 0, nil
}
