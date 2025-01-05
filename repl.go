package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Mooogy/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func initRepl() {
	cfg := &config{pokeapiClient: pokeapi.NewClient(5 * time.Second, 30 * time.Second)}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		validCommands := getCommands()
		if command, ok := validCommands[commandName]; ok {
			err := command.callback(cfg)
			if err != nil {
				fmt.Printf("\n%v\n\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Use \"help\" for a list of valid commands")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	clean := strings.Fields(lowered)
	return clean
}

type cliCommand struct {
	name			string
	description		string
	callback		func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a list of available commands",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Closes the PokedexCLI",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Displays the next page of location-areas",
			callback:  commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous page of location-areas",
			callback:  commandMapb,
		},
	}
}