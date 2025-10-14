package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreaList struct {
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []namedAPIResource `json:"results"`
}

type namedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(cfg *config) error {
	url := cfg.nextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations locationAreaList

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}

	if locations.Next != nil {
		cfg.nextURL = *locations.Next
	} else {
		cfg.nextURL = ""
	}

	if locations.Previous != nil {
		cfg.prevURL = *locations.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := cfg.prevURL
	if url == "" {
		fmt.Println("You are on the first page")
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations locationAreaList

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	for _, r := range locations.Results {
		fmt.Println(r.Name)
	}

	if locations.Next != nil {
		cfg.nextURL = *locations.Next
	} else {
		cfg.nextURL = ""
	}

	if locations.Previous != nil {
		cfg.prevURL = *locations.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}
