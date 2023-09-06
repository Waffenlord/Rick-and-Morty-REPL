package main

import (
	"testing"
)

// Test the cleanInput function
func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WORLD",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "1 2 3",
			expected: []string{
				"1",
				"2",
				"3",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		// Compare the length of the two arrays
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}

		// Compare individual characters of each array
		for i := range actual {
			actualChar := actual[i]
			expectedChar := cs.expected[i]
			if actualChar != expectedChar {
				t.Errorf("The characters are not equal: %v vs %vs", actualChar, expectedChar)
			
			}

		}
	}
}

// Test the getSecondaryCommands function
func TestGetSecondaryCommands(t *testing.T) {
	cases := []struct{
		input string
		expected map[string]cliCommand
	}{
		{
			input: "character",
			expected:  map[string]cliCommand{
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
					description: "Type this command follow by the id of the character you want to see.",
					callback: commandViewChar,
				},
			},
		},
	}

	for _, cs := range cases {
		actual := getSecondaryCommands(cs.input)
		
		// Compare the length of the maps
		if len(actual) != len(cs.expected) {
			t.Errorf("The number of commands is inconsistent: %v vs %v", len(actual), len(cs.expected))
		}

		// Compare keys and values of the maps
		for key, expectedValue := range cs.expected {
			actualvalue, ok := actual[key]
			if !ok {
				t.Errorf("Missing command: %s", key)
			}

			if actualvalue.name != expectedValue.name ||
				actualvalue.description != expectedValue.description {
					t.Errorf("Values are not equal: %v vs %v", actualvalue, expectedValue)
				}


		}
	}
}

