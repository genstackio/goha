package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func UpdatePartnerNotificationUrl(args []string, env string) (int, error) {
	if len(args) < 3 {
		return 1, errors.New("<clientId> <clientSecret> <url>")
	}
	clientId := args[0]
	clientSecret := args[1]
	c := goha.Client{}
	c.Init(clientId, clientSecret, env)
	r, err := c.UpdatePartnerNotificationUrl(args[2])
	if err != nil {
		return 5, err
	}
	raw, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return 6, err
	}
	fmt.Println(string(raw))

	return 0, nil
}
