package main

import (
	"time"

	"github.com/ArashPoorazam/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(20*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
