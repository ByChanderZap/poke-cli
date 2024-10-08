package main

import "github.com/ByChanderZap/poke-cli/internal/pokeapi"

type config struct {
	pokeApiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
