package main

import (
	"time"

	"github.com/BytteM3/pokedex/internal/pokecache"
)

func main() {
	cfg := &config{}
	c := pokecache.NewCache(time.Duration(5 * time.Second))
	p := make(map[string]PokemonData)
	cfg.cache = c
	cfg.pokedex = p
	startRepl(cfg)
}
