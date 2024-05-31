package pokeapi

type ExploredPokemons struct {
	ID                int `json:"id"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
