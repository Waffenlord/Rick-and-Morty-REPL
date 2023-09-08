package main

import "fmt"

// Return to the main menu from the character REPL
func commandCharBack(cfg *config, args ...string) error {
	fmt.Println("Returning to the main menu...")
	return nil
}