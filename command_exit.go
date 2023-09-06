package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Terminating the program...")
	os.Exit(0)
	return nil
}