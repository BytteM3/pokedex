package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %v\n", userInput[0])
	}
}
