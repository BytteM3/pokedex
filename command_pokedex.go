package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Printf("Your Pokedex:\n")
	for pokemon := range cfg.pokedex {
		fmt.Printf("- %v\n", pokemon)
	}
	return nil
}
