package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(name string) (Pokemon, error){
	url := baseURL + "/pokemon/" + name

	// Check cache
	if data, ok := c.cache.Get(url); ok {
		var pokemon Pokemon
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// Get pokemon data
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}