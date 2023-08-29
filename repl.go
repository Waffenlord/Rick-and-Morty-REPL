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

		err := command.callback(c)
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
	callback func(*config) error
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
		"character": {
			name: "character",
			description: "Access all characters information.",
			callback: commandCharRepl,
		},
	}
}

func getSecondaryCommands(option string) map[string]cliCommand {
	commandsMap := make(map[string]map[string]cliCommand)
	commandsMap["character"] = map[string]cliCommand{"help": {
		name: "help",
		description: "See the list of available commands for characters",
		callback: commandCharHelp,
		},
		"back": {
			name: "back",
			description: "Go back to the main menu.",
			callback: commandCharBack,
		},
		"map": {
			name: "map",
			description: "Get the a list with the next 20 characters",
			callback: commandCharMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get the a list with the previous 20 characters",
			callback: commandCharMapB,
		},
		"exit": {
			name: "exit",
			description: "Exit the program.",
			callback: commandExit,
		},
	}
	
	currentCommands, ok := commandsMap[option]
	if !ok {
		fmt.Println("can't retrieve the commands!")
		return nil
	}
	return currentCommands
}