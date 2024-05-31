package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pokemon struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatchPokemon(c *config) error {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the name of Pokemon > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSpace(strings.ToLower(input))
		findPokemon(input)
		break
	}
	return nil
}

func findPokemon(pokemonName string) error {
	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	var pokemon Pokemon
	err := getJSON(pokemonURL, &pokemon)
	if err != nil {
		fmt.Println("Error fetching Pokémon details:", err)
		return nil
	}

	fmt.Printf("Name: %s\nID: %d\nTypes:\n", pokemon.Name, pokemon.ID)
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}
