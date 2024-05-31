package pokeapi

type Pokemon struct {
	Name    string `json:"name"`
	ID      int    `json:"id"`
	BaseExp int    `json:"base_experience"`
	Weight  int    `json:"weight"`
	Height  int    `json:"height"`
	Types   []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
