package main

import (
	"fmt"
	"os"
)

func commandExit(*config) error {
	fmt.Println("Closing PokedexCLI... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}