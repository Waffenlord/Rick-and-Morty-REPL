package main

import "fmt"

func commandCharHelp(cfg *config) error {
	fmt.Println("---------------------------------")
	fmt.Println("These are the available character commands:")
	availableCommands := getSecondaryCommands("character")
	for _, command := range availableCommands {
		fmt.Printf("%s : %s \n", command.name, command.description)
	}
	fmt.Println("---------------------------------")

	return nil
}