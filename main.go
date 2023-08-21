package main

import rick_morty_api "github.com/Waffenlord/Rick-and-Morty-REPL/internal"

type config struct {
	rickMortyClient rick_morty_api.Client
}

func main() {

	cfg := config{
		rickMortyClient: rick_morty_api.NewClient(),
	}

	startRepl(&cfg)
}
