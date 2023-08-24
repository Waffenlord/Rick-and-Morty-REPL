package main

import "fmt"

func commandCharMap(cfg *config) error {
	resp, err := cfg.rickMortyClient.ListCharacters()
	if err != nil {
		return err
	}

	fmt.Println("Characters: ")
	for _, character := range resp.Results {
		fmt.Printf("Name: %v, Status: %v \n", character.Name, character.Status)
	}

	return nil
}