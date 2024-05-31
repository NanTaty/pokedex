package main

import "fmt"

func commandHelp(c *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println(" ")
	for _, cli := range getCommands() {
		fmt.Printf("%s: %s \n", cli.name, cli.description)
	}
	return nil
}
