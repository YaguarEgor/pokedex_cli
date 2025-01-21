package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YaguarEgor/caching"
)

const baseURL = "https://pokeapi.co/api/v2/"

type PokeLocationJSON struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
	fmt.Println("Hi2")
	cache.Add(cur_url, bytes)
	fmt.Println("Hi")
	return data, nil
}
