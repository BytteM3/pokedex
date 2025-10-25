package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	pokemon, exists := cfg.pokedex[args[0]]
	if !exists {
		fmt.Printf("%v has never been caught!\n", args[0])
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, value := range pokemon.Stats {
			fmt.Printf(" -%v: %v\n", value.Stat.StatName, value.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, value := range pokemon.Types {
			fmt.Printf(" - %v\n", value.Type.TypeName)
		}
	}

	return nil
}
