package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	pokeApiClient := cfg.pokeApiClient
	res, err := pokeApiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, l := range res.Results {
		fmt.Printf(" - %s\n", l.Name)
	}
	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaUrl == nil {
		return errors.New("you're on the first page")
	}
	pokeApiClient := cfg.pokeApiClient
	res, err := pokeApiClient.ListLocationAreas(cfg.previousLocationAreaUrl)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, l := range res.Results {
		fmt.Printf(" - %s\n", l.Name)
	}
	cfg.nextLocationAreaUrl = res.Next
	cfg.previousLocationAreaUrl = res.Previous
	return nil
}
