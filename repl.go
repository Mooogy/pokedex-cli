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
	pokedex map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func initRepl() {
	cfg := &config{pokeapiClient: pokeapi.NewClient(5 * time.Second, 90 * time.Second), pokedex: map[string]pokeapi.Pokemon{}}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		args := input[1:]
		validCommands := getCommands()
		if command, ok := validCommands[commandName]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Printf("%v\n\n", err)
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
	callback		func(*config, ...string) error
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
			description: "Displays the next page of locations",
			callback:  commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous page of locations",
			callback:  commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Displays all encounterable pokemon at the named location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Attempts to catch the named pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "Displays information about pokemon you caught",
			callback: commandInspect,
		},
	}
}