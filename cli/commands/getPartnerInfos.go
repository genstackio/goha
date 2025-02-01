package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/genstackio/goha"
)

func GetPartnerInfos(args []string, env string) (int, error) {
	if len(args) < 2 {
		return 1, errors.New("<clientId> <clientSecret>")
	}
	clientId := args[0]
	clientSecret := args[1]
	c := goha.Client{}
	c.Init(clientId, clientSecret, env)
	infos, err := c.GetPartnerInfos()
	if err != nil {
		return 5, err
	}
	raw, err := json.MarshalIndent(infos, "", "  ")
	if err != nil {
		return 6, err
	}
	fmt.Println(string(raw))

	return 0, nil
}
