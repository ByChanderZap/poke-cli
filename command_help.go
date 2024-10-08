package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Poke-CLI")
	fmt.Println("Here are the available commands:")

	availableCommands := getCommands()

	for _, ac := range availableCommands {
		fmt.Printf(" - %v: %v\n", ac.name, ac.description)
	}

	return nil
}
