package main

import (
	"bufio"
	"fmt"
	"os"
)

// Loop that evaluates character specific commands
func commandCharRepl(cfg *config, args ...string) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter a command to explore charaters > ")

		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		args := []string{}
		if len(cleanedInput) == 0 {
			continue
		}

		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		commandName := cleanedInput[0]
		availableCommands := getSecondaryCommands("character")

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command!")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

		if command.name == "back"{
			break
		}

	}

	return nil
}



