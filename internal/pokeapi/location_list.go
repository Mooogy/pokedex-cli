package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c* Client) GetLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	var data LocationResponse
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return LocationResponse{}, err
	}

	return data, nil
}