package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandCharRepl(cfg *config) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter a command to explore charaters > ")

		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		availableCommands := getSecondaryCommands("character")

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command!")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

		if command.name == "back"{
			break
		}

	}

	return nil
}



