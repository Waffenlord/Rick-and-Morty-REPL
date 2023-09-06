package main

import (
	"time"

	rick_morty_api "github.com/Waffenlord/Rick-and-Morty-REPL/internal"
)

type config struct {
	rickMortyClient rick_morty_api.Client
	nextCharacterList *string
	previousCharacterList *string
	characterCount int
}

func main() {

	cfg := config{
		rickMortyClient: rick_morty_api.NewClient(time.Hour),
		characterCount: 826,
	}

	startRepl(&cfg)
}
