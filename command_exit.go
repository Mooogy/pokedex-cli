package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing PokedexCLI... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}