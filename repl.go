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

		command := input[0]
		fmt.Printf("Your command was: %s\n", command)
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	clean := strings.Fields(lowered)
	return clean
}