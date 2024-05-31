package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Nan_taty/pokedex/pokeapi"
	"github.com/Nan_taty/pokedex/pokecache"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokeCache        *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(c *config) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("pokedex > ")
		commandName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		commandName = strings.TrimSpace(strings.ToLower(commandName))
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c)
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("Command not found")
		}
		continue
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			name:        "catch",
			description: "tries to catch pokemon",
			callback:    commandCatchPokemon,
		},
		"explore": {
			name:        "explore",
			description: "explores pokemons in area",
			callback:    commandExplore,
		},
	}
}

func getJSON(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	name, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return json.Unmarshal(name, target)
}
