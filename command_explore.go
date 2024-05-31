package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandExplore(cfg *config) error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the name of Location > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSpace(strings.ToLower(input))
		exploredPokemons, err := cfg.pokeapiClient.ExploreLocation(input)
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
