package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter some command > ")

		scanner.Scan()
		text := scanner.Text()

		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command!")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}

	}
}



func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)

	return words
}

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message.",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the program.",
			callback: commandExit,
		},
	}
}