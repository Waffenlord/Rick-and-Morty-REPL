package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandCharMap(cfg *config, args ...string) error {
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

func commandCharMapB(cfg *config, args ...string) error {
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

func commandViewChar(cfg *config, args ...string) error {
	charId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	
	if charId < 1 || charId > cfg.characterCount {
		return errors.New("invalid id! out of range.")
	}

	resp, err := cfg.rickMortyClient.GetSingleCharacter(fmt.Sprintf("%v", charId))
	if err != nil {
		return err
	}

	fmt.Println("==========================================")
	fmt.Printf("Printing information of: %s\n", resp.Name)
	fmt.Println("==========================================")
	fmt.Printf(" - ID: %v\n", resp.ID)
	fmt.Printf(" - Status: %v\n", resp.Status)
	fmt.Printf(" - Species: %v\n", resp.Species)


	return nil
}