package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Print("Welcome to PokedexCLI!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}