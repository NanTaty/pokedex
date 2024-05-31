package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(locationName string) (ExploredPokemons, error) {
	url := baseURL + "/location-area/" + locationName
	if locationName == "" {
		return ExploredPokemons{}, errors.New("No string provided")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := ExploredPokemons{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return ExploredPokemons{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploredPokemons{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploredPokemons{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploredPokemons{}, err
	}

	pokemonResp := ExploredPokemons{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return ExploredPokemons{}, err
	}

	return pokemonResp, nil
}
