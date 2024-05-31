package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	for {
		if len(args) != 1 {
			return errors.New("you must provide a location name")
		}
		name := args[0]
		exploredPokemons, err := cfg.pokeapiClient.ExploreLocation(name)
		if err != nil {
			return err
		}

		fmt.Println("Found Pokemon:")
		for _, location := range exploredPokemons.PokemonEncounters {
			fmt.Printf("- %s \n", location.Pokemon.Name)
		}
		break
	}
	return nil
}
