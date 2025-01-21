package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/YaguarEgor/caching"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config) error
}

func startRepl() {
	commands := getAllCommands()
	config := Config{
		NextLocation: nil,
		PreviousLocation: nil,
		cache: caching.NewCache(time.Second*10),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pockedex > ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("\nGoodbye")
			break
		}
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		command, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&config)
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func getAllCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "Exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: "Map",
			description: "Show next 20 location areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "MapB",
			description: "Show previous 20 location areas in the Pokemon world",
			callback: commandMapB,
		},
		"explore": {
			name: "Explore",
			description: "Get all Pokemons here",
			callback: commandExplore,
		},
	}
}
