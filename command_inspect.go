package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("pokemon must be catch first")
	}

	fmt.Printf("Name:\n - %s\n", pokemon.Name)
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	return nil
}
