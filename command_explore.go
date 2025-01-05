package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a location name")
	}

	name := args[0]
	encounters, err := cfg.pokeapiClient.GetEncountersByLocation(name)
	if err != nil {
		return fmt.Errorf("explore error: %v", err)
	}

	fmt.Println("Exploring " + encounters.Name + "...")
	for _, mon := range encounters.Pokemon_encounters {
		fmt.Println(mon.Pokemon.Name)
	}
	fmt.Println()

	return nil
}