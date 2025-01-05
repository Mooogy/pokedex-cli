package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.pokedex {
		fmt.Printf("\t- %s\n", k)
	}
	fmt.Println()
	return nil
}