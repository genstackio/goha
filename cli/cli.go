package main

import (
	"fmt"
	"github.com/genstackio/goha"
	"os"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Syntax: goha <clientId> <clientSecret> <env> <action>")
		os.Exit(1)
	}
	clientId := os.Args[1]
	clientSecret := os.Args[2]
	env := os.Args[3]
	action := os.Args[4]
	switch action {
	case "list-organizations":
		c := goha.Client{}
		c.Init(clientId, clientSecret, env)
		p, err := c.GetMyOrganizations()
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(4)
		}
		for _, v := range p.Items {
			fmt.Println(v.Name, v.City, v.OrganizationSlug, v.Url)
		}
	case "list-organization-payments":
		if len(os.Args) < 6 {
			fmt.Println("Syntax: goha <clientId> <clientSecret> <env> list-organization-payments <organizationSlug>")
			os.Exit(1)
		}
		c := goha.Client{}
		c.Init(clientId, clientSecret, env)
		p, err := c.GetOrganizationPayments(os.Args[5], goha.GetPaymentsOptions{})
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			os.Exit(5)
		}
		for _, v := range p.Items {
			fmt.Println(v.Id, v.Amount, v.Date)
		}
	case "test-credentials":
		ok, err := goha.TestApiCredentials(clientId, clientSecret, env)

		if !ok {
			fmt.Println("NOK - Bad credentials:" + err.Error())
			os.Exit(3)
		}
		fmt.Println("OK")
	default:
		fmt.Println("Error: unrecognized action '" + action + "'")
		os.Exit(2)
	}
}
