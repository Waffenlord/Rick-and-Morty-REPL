package main

import (
	"time"

	rick_morty_api "github.com/Waffenlord/Rick-and-Morty-REPL/internal"
)

type config struct {
	rickMortyClient rick_morty_api.Client
	nextCharacterList *string
	previousCharacterList *string
}

func main() {

	cfg := config{
		rickMortyClient: rick_morty_api.NewClient(time.Hour),
	}

	startRepl(&cfg)
}
