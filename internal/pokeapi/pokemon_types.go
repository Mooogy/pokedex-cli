package pokeapi

type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Stats []struct {
		Base_stat int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}