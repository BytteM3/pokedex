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

func commandMap(cfg *config, args []string) error {
	url := cfg.nextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	data, ok := cfg.cache.Get(url)
	if !ok {
		fmt.Println("------using network------")
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

	var locations locationAreaList

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

func commandMapb(cfg *config, args []string) error {
	url := cfg.prevURL
	if url == "" {
		fmt.Println("You are on the first page")
	}

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

	var locations locationAreaList

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
