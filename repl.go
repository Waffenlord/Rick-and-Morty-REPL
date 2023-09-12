package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// Main loop of the REPL
func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Please enter some command > ")

		scanner.Scan()
		text := scanner.Text()

		// Clean user input
		cleanedInput := cleanInput(text)
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		// Get the main menu commands
		availableCommands := getCommands()

		// Check if the user input is a valid command
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command!")
			continue
		}

		// Call the selected command
		err := command.callback(c)
		if err != nil {
			fmt.Println(err)
		}

	}
}


// Clean user input by setting it to lower case and extracting individual words
func cleanInput(input string) []string {
	lowerCase := strings.ToLower(input)
	words := strings.Fields(lowerCase)

	return words
}

//Structure of every command in the REPL
type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

// Return the main menu commands
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

// Return the each section specific commands
func getSecondaryCommands(option string) map[string]cliCommand {
	commandsMap := make(map[string]map[string]cliCommand)
	commandsMap["character"] = map[string]cliCommand{
		"help": {
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
		"view": {
			name: "view",
			description: "Type this command followed by the id of the character you want to see.",
			callback: commandViewChar,
		},
		"save": {
			name: "save",
			description: "Type this command followed by the id of the character you want to save.",
			callback: commandSaveChar,
		},
		"list": {
			name: "list",
			description: "See your saved characters",
			callback: commandListChars,
		},
		"del": {
			name: "del",
			description: "Type this command folloed by the id of the character you want to delete",
			callback: commandDeleteChar,
		},
		"export": {
			name: "export",
			description: "Export your saved character",
			callback: commandExportChars,
		},
	}
	
	currentCommands, ok := commandsMap[option]
	if !ok {
		fmt.Println("can't retrieve the commands!")
		return nil
	}
	return currentCommands
}