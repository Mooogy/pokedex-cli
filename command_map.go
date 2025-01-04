package main

import "fmt"

func commandMapf(cfg *config) error {
	// Use config's nextLocationsURL, GetLocations will use first page URL if nil
	locationsRes, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	// Set config
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	// List locations
	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}

func commandMapb(cfg *config) error {
	// Check if on first page
	if cfg.prevLocationsURL == nil {
		return fmt.Errorf("already on the first page")
	}

	// Use config's nextLocationsURL, GetLocations will use first page URL if nil
	locationsRes, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	// Set config
	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	// List locations
	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return nil
}