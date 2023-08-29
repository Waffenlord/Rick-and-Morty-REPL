package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Terminating the program...")
	os.Exit(0)
	return nil
}