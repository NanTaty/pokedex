package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.userPokedex) == 0 {
		fmt.Println("You haven't caught any pokemon")
	}
	fmt.Println("Your pokedex:")
	for _, userPokedex := range c.userPokedex {
		fmt.Printf("- %s \n", userPokedex.Name)
	}
	return nil
}
