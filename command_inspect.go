package main

import (
	"errors"
	"fmt"
)

func commandInspectPokemon(c *config, args ...string) error {
	for {
		if len(args) != 1 {
			return errors.New("You must provide pokemon to inspect!")
		}
		name := args[0]
		pokemon, exists := c.userPokedex[name]
		if !exists {
			fmt.Println("you have not caught that pokemon")
			return nil
		}

		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %v \n", pokemon.Height)
		fmt.Printf("Weight: %v \n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %v \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("- %s\n", pokemonType.Type.Name)
		}
		break
	}
	return nil
}
