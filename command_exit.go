package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing PokedexCLI... Goodbye!")
	os.Exit(0)
	return nil
}