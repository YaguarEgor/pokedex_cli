package main

import (
	"fmt"
	"github.com/YaguarEgor/caching"
	"os"
	
)

type Config struct {
	NextLocation *string
	PreviousLocation *string
	cache *caching.Cache
}

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getAllCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *Config) error {
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

func commandMapB(conf *Config) error {
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

func commandExplore(conf *Config) error {
	return nil
}