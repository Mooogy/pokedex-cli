package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c* Client) GetLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if entry, ok := c.cache.Get(url); ok {
		var locations LocationResponse
		err := json.Unmarshal(entry, &locations)
		if err != nil {
			return LocationResponse{}, err
		}
		fmt.Println("DATA PULLED FROM CACHE")
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	c.cache.Add(url, data)

	var locations LocationResponse
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationResponse{}, err
	}

	return locations, nil
}