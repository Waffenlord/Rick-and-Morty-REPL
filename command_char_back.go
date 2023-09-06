package main

import "fmt"

func commandCharBack(cfg *config, args ...string) error {
	fmt.Println("Returning to the main menu...")
	return nil
}