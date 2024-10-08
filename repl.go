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

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command.")
			continue
		}

		err := command.callback(cfg)
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
		"exit": {
			name:        "exit",
			description: "Closes the Poke-CLI",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List all locations in the pokeworld",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previos page of locations in the pokeworld",
			callback:    callbackMapb,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(s string) []string {
	lowered := strings.ToLower(s)
	splitted := strings.Fields(lowered)
	return splitted
}
