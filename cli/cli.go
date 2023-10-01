package main

import (
	"fmt"
	"github.com/genstackio/goha/cli/commands"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Syntax: goha <env> <action>")
		os.Exit(1)
	}
	env := os.Args[1]
	action := os.Args[2]
	cmd, ok := commands.Commands[action]
	if !ok {
		fmt.Println("Error: unrecognized action '" + action + "'")
		os.Exit(2)
	}
	code, err := cmd(os.Args[3:], env)

	if 0 != code {
		if err != nil {
			if 1 == code {
				fmt.Println("Syntax: goha " + env + " " + action + " " + err.Error())
			} else {
				fmt.Println("Error: " + err.Error())
			}
		}
		os.Exit(code)
	}
}
