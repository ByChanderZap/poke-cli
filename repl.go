package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Poke-CLI> ")
		scanner.Scan()
		data := scanner.Text()
		c := cleanInput(data)
		if len(c) == 0 {
			continue
		}
		commandName := c[0]

		args := []string{}
		if len(c) > 1 {
			args = c[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show available commands",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List all locations in the pokeworld",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous page of locations in the pokeworld",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List the pokemons that appears on a specific area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch a pokemon and add it to your pokedex!",
			callback:    callbackCatch,
		},
		"exit": {
			name:        "exit",
			description: "Closes the Poke-CLI",
			callback:    callbackExit,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	splitted := strings.Fields(lowered)
	return splitted
}
