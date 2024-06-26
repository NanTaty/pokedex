package main

import (
	"github.com/Nan_taty/pokedex/pokeapi"
	"time"
)

func main() {
	// Initialize the cache
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		userPokedex:   map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
