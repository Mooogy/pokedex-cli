package pokeapi

type LocationResponse struct {
	Next 		*string 	`json:"next"`
	Previous 	*string 	`json:"previous"`
	Results []struct{
		Name 	string 		`json:"name"`
		Url 	string 		`json:"url"`
	} `json:"results"`
}

type EncountersResponse struct {
	Name		string 		`json:"name"`
	Pokemon_encounters []struct {
		Pokemon	struct {
			Name	string 	`json:"name"`
			Url		string	`json:"url"`
		}
	} `json:"pokemon_encounters"`
}