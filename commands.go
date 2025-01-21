package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/YaguarEgor/caching"
)

type Config struct {
	NextLocation     *string
	PreviousLocation *string
	cache            *caching.Cache
	pokedex 		 map[string]Pokemon
}

func commandExit(conf *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getAllCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *Config, args ...string) error {
	data, err := GetLocations(conf.NextLocation, conf.cache)
	if err != nil {
		return fmt.Errorf("error in map: %v", err)

	}
	conf.PreviousLocation = data.Previous
	conf.NextLocation = data.Next
	for _, el := range data.Results {
		fmt.Println(el.Name)
	}
	return nil
}

func commandMapB(conf *Config, args ...string) error {
	data, err := GetLocations(conf.PreviousLocation, conf.cache)
	if err != nil {
		return fmt.Errorf("error in mapb: %v", err)
	}
	conf.PreviousLocation = data.Previous
	conf.NextLocation = data.Next
	for _, el := range data.Results {
		fmt.Println(el.Name)
	}
	return nil
}

func commandExplore(conf *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to print name of location to get pokemons through 'explore'")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	location, err := getPokemons(conf.NextLocation, conf.cache, args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon: \n")
	for _, val := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", val.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to print name of pokemon to 'catch' him")
	}
	pokemon, err := getOnePokemon(conf.cache, args[0])
	if err != nil {
		return err
	}
	
	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	conf.pokedex[pokemon.Name] = pokemon
	return nil
}

func commandInspect(conf *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you need to print name of pokemon to 'inspect' him")
	}
	pokemon, ok := conf.pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} 
	
	fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
`, pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, el := range pokemon.Types {
		fmt.Printf("  -%s\n", el.Type.Name)
	}
	return nil
}

func commandPokedex(conf *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range conf.pokedex {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}