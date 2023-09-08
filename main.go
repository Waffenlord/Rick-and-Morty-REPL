package main

import (
	"time"

	rick_morty_api "github.com/Waffenlord/Rick-and-Morty-REPL/internal"
)

// Configuration struct
type config struct {
	rickMortyClient rick_morty_api.Client
	nextCharacterList *string
	previousCharacterList *string
	characterCount int
}

func main() {
	//Create an instance of the config struct
	cfg := config{
		rickMortyClient: rick_morty_api.NewClient(time.Hour),
		characterCount: 826,
	}

	//Pass a pointer of the config struct to the main REPL
	startRepl(&cfg)
}
