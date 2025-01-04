package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokemon > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		commandName := input[0]
		validCommands := getCommands()
		if command, ok := validCommands[commandName]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
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
	callback		func() error
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
	}
}