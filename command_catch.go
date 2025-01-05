package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please enter a valid pokemon name")
	}

	pokemonName := args[0]
	
	// Get pokemon data 
	pokemon, err := cfg.pokeapiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		return nil
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	
	// Use pokemon base experience to determine catch rate (608 is highest (BLISSEY))
	// Custom formula: CR = 20 + 80 * (1 - (baseEXP / 608))
	catchRate := int32(20.0 + 75.0 * (1.0 - (float64(pokemon.BaseExperience) / 608.0)))
	if caught := rand.Int31n(101) <= int32(catchRate); caught {
		fmt.Println(pokemonName + " was caught!\n")
		cfg.pokedex[pokemonName] = pokemon
		return nil
	}

	fmt.Println(pokemonName + " escaped!\n")
	return nil
}