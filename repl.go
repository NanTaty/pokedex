package main

import (
	"bufio"
	"fmt"
	"github.com/Nan_taty/pokedex/pokeapi"
	"github.com/Nan_taty/pokedex/pokecache"
	"os"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
	userPokedex      map[string]pokeapi.Pokemon
}

func startRepl(c *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations.",
			callback:    commandMapb,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "tries to catch pokemon",
			callback:    commandCatchPokemon,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "explores pokemons in area",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspects pokemon in user's pokedex",
			callback:    commandInspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "displays all caught pokemon",
			callback:    commandPokedex,
		},
	}
}
