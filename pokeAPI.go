package main

import (
	"encoding/json"
	"net/http"

	"github.com/YaguarEgor/caching"
)

const baseURL = "https://pokeapi.co/api/v2/"

func GetLocations(nextUrl *string, cache *caching.Cache) (PokeLocationJSON, error) {
	var data PokeLocationJSON
	cur_url := baseURL + "location-area/"
	if nextUrl != nil {
		cur_url = *nextUrl
	}

	val, ok := cache.Get(cur_url)
	if ok {
		err := json.Unmarshal(val, &data)
		if err != nil {
			return PokeLocationJSON{}, err
		}
		return data, nil
	}
	req, err := http.NewRequest("GET", cur_url, nil)
	if err != nil {

		return PokeLocationJSON{}, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return PokeLocationJSON{}, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return PokeLocationJSON{}, err
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return PokeLocationJSON{}, err
	}
	cache.Add(cur_url, bytes)
	return data, nil
}

func getPokemons(nextUrl *string, cache *caching.Cache, areaName string) (Location, error) {
	cur_url := baseURL + "location-area/"
	if nextUrl != nil {
		cur_url = *nextUrl
	}
	cur_url += areaName
	var pokemons Location
	val, ok := cache.Get(cur_url)
	if ok {
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return Location{}, err
		}
		return pokemons, nil
	}
	req, err := http.NewRequest("GET", cur_url, nil)
	if err != nil {
		return Location{}, err
	}
	client := &http.Client{}
	data, err := client.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer data.Body.Close()
	decoder := json.NewDecoder(data.Body)
	err = decoder.Decode(&pokemons)
	if err != nil {
		return Location{}, err
	}
	bytes, err := json.Marshal(pokemons)
	if err != nil {
		return Location{}, err
	}
	cache.Add(cur_url, bytes)
	return pokemons, nil
}

func getOnePokemon(cache *caching.Cache, pokemonName string) (Pokemon, error) {
	var pokemon Pokemon
	if val, ok := cache.Get(pokemonName); ok {
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}
	cur_url := baseURL + "/pokemon/" + pokemonName
	req, err := http.NewRequest("GET", cur_url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	bytes, err := json.Marshal(pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(cur_url, bytes)
	return pokemon, nil
}
