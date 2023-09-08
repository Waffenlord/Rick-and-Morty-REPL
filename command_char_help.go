package main

import "fmt"

// Shows the available commands specific to characters
func commandCharHelp(cfg *config, args ...string) error {
	fmt.Println("---------------------------------")
	fmt.Println("These are the available character commands:")
	availableCommands := getSecondaryCommands("character")
	for _, command := range availableCommands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	fmt.Println("---------------------------------")

	return nil
}