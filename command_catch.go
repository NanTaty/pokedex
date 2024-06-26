package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatchPokemon(c *config, args ...string) error {
	for {
		if len(args) != 1 {
			return errors.New("you must provide a pokemon name")
		}

		name := args[0]
		pokemon, err := c.pokeapiClient.GetPokemonByName(name)
		if err != nil {
			return err
		}

		res := rand.Intn(pokemon.BaseExp)

		fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
		if res > 40 {
			fmt.Printf("%s escaped!\n", pokemon.Name)
			return nil
		}

		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command")
		c.userPokedex[pokemon.Name] = pokemon

		break
	}
	return nil
}
