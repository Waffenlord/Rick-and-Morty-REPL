package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/Waffenlord/Rick-and-Morty-REPL/internal"
)

// Returns the next 20 characters
func commandCharMap(cfg *config, args ...string) error {
	// Call the ListCharacters method to make the request
	resp, err := cfg.rickMortyClient.ListCharacters(cfg.nextCharacterList)
	if err != nil {
		return err
	}

	// Print the returned list of characters
	fmt.Println("Characters: ")
	for _, character := range resp.Results {
		fmt.Printf("ID: %v, Name: %v, Status: %v \n", character.ID, character.Name, character.Status)
	}

	//Store the next character list URL in the config struct 
	cfg.nextCharacterList = resp.Info.Next
	//Store the previous character list URL in the config struct 
	cfg.previousCharacterList = resp.Info.Prev
	return nil
}

// Returns the previous 20 characters
func commandCharMapB(cfg *config, args ...string) error {
	// Check if it is the first group of characters
	if cfg.previousCharacterList == nil {
		return errors.New("you are on the first page.")
	}

	// Call the ListCharacters method to make the request
	resp, err := cfg.rickMortyClient.ListCharacters(cfg.previousCharacterList)
	if err != nil {
		return err
	}

	// Print the returned list of characters
	fmt.Println("Characters: ")
	for _, character := range resp.Results {
		fmt.Printf("ID: %v, Name: %v, Status: %v \n", character.ID, character.Name, character.Status)
	}

	//Store the next character list URL in the config struct 
	cfg.nextCharacterList = resp.Info.Next
	//Store the previous character list URL in the config struct 
	cfg.previousCharacterList = resp.Info.Prev
	return nil
}

// Shows information about an specific character
func commandViewChar(cfg *config, args ...string) error {
	//Check if the first argument is a number
	charId, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	
	//Check if the ID is within the number of available characters
	if charId < 1 || charId > cfg.characterCount {
		return errors.New("invalid id! out of range.")
	}

	// Call the GetSingleCharacter to make the request
	resp, err := cfg.rickMortyClient.GetSingleCharacter(fmt.Sprintf("%v", charId))
	if err != nil {
		return err
	}

	// Print the character's information
	fmt.Println("==========================================")
	fmt.Printf("Printing information of: %s\n", resp.Name)
	fmt.Println("==========================================")
	fmt.Printf(" - ID: %v\n", resp.ID)
	fmt.Printf(" - Status: %v\n", resp.Status)
	fmt.Printf(" - Species: %v\n", resp.Species)
	fmt.Printf(" - Gender: %v\n", resp.Gender)
	fmt.Printf(" - Origin: %v\n", resp.Origin.Name)
	fmt.Printf(" - Location: %v\n", resp.Location.Name)

	return nil
}

// Saves a specific character inside the Client instance
func commandSaveChar(cfg *config, args ...string) error {
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

	char := rick_morty_api.SavedChar{
		ID: resp.ID,
		Name: resp.Name,
		Status: resp.Status,
		Species: resp.Species,
		Type: resp.Type,
		Gender: resp.Gender,
		Origin: resp.Origin.Name,
		Location: resp.Location.Name,
		Image: resp.Image,
	}
	fmt.Printf("Saving character...\n")

	// Add an entry to the SavedChars struct
	cfg.rickMortyClient.SavedChars.Characters[resp.ID] = char

	fmt.Printf("ID: %v, Name: %s saved!\n", resp.ID, resp.Name)

	return nil

}