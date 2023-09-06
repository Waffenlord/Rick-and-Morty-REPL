package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("---------------------------------")
	fmt.Println("These are the available commands:")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	fmt.Println("---------------------------------")

	return nil
}