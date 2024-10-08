package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config) error {
	fmt.Println("Closing Poke-CLI")
	os.Exit(0)
	return nil
}
