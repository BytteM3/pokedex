package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location-area-name>")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	fmt.Printf("Exploring %s...\n", args[0])

	data, ok := cfg.cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", res.Status)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cfg.cache.Add(url, body)
		data = body
	}

	var locationPokemon LocationArea

	if err := json.Unmarshal(data, &locationPokemon); err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")

	for _, pokemonEncounter := range locationPokemon.PokemonEncounters {
		fmt.Printf("- %s\n", pokemonEncounter.Pokemon.Name)
	}
	return nil

}
