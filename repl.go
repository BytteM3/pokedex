package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextURL string
	prevURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
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
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		command, exists := commands[userInput[0]]
		if exists {
			command.callback(cfg)
		} else {
			fmt.Printf("Unknown command\n")
		}
	}
}
