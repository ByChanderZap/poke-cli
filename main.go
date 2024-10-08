package main

import (
	"time"

	"github.com/ByChanderZap/poke-cli/internal/pokeapi"
)

type config struct {
	pokeApiClient           pokeapi.Client
	nextLocationAreaUrl     *string
	previousLocationAreaUrl *string
	caughtPokemon           map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeApiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
