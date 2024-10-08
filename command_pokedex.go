package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

	pokemons := cfg.caughtPokemon
	fmt.Println("Pokemon in your pokedex:")
	for _, poke := range pokemons {
		fmt.Printf(" ~ %s", poke.Name)
	}

	return nil
}
