package main

import (
	"errors"
	"fmt"
)

func commandCharMap(cfg *config) error {
	resp, err := cfg.rickMortyClient.ListCharacters(cfg.nextCharacterList)
	if err != nil {
		return err
	}

	fmt.Println("Characters: ")
	for _, character := range resp.Results {
		fmt.Printf("ID: %v, Name: %v, Status: %v \n", character.ID, character.Name, character.Status)
	}

	cfg.nextCharacterList = resp.Info.Next
	cfg.previousCharacterList = resp.Info.Prev
	return nil
}

func commandCharMapB(cfg *config) error {
	if cfg.previousCharacterList == nil {
		return errors.New("you are on the first page.")
	}
	resp, err := cfg.rickMortyClient.ListCharacters(cfg.previousCharacterList)
	if err != nil {
		return err
	}

	fmt.Println("Characters: ")
	for _, character := range resp.Results {
		fmt.Printf("ID: %v, Name: %v, Status: %v \n", character.ID, character.Name, character.Status)
	}

	cfg.nextCharacterList = resp.Info.Next
	cfg.previousCharacterList = resp.Info.Prev
	return nil
}