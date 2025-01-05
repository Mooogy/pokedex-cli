package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c* Client) GetEncountersByLocation(location string) (EncountersResponse, error) {
	// Location area path + <location-area>
	url := baseURL + "/location-area/" + location

	// Check if data is in cache
	if entry, ok := c.cache.Get(url); ok {
		fmt.Println("PULLING FROM CACHE")
		var encounters EncountersResponse
		err := json.Unmarshal(entry, &encounters)
		if err != nil {
			return EncountersResponse{}, err
		}
		return encounters, nil
	}

	// Grab data using pokeapi
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return EncountersResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return EncountersResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EncountersResponse{}, err
	}

	// Cache the data for subsequent use
	c.cache.Add(url, data)

	var encounters EncountersResponse
	err = json.Unmarshal(data, &encounters)
	if err != nil {
		return EncountersResponse{}, err
	}

	return encounters, nil
}