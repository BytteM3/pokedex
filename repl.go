package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BytteM3/pokedex/internal/pokecache"
)

type config struct {
	nextURL string
	prevURL string
	cache   *pokecache.Cache
	pokedex map[string]PokemonData
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
}

var commands map[string]cliCommand

func cleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}

func startRepl(cfg *config) error {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exits the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays 20 locations, use again for next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explores the area provided as argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "tries to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists all caught pokemon",
			callback:    commandPokedex,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		tokens := cleanInput(scanner.Text())
		name := tokens[0]
		args := tokens[1:]
		if len(tokens) == 0 {
			continue
		}
		command, exists := commands[name]
		if exists {
			command.callback(cfg, args)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
