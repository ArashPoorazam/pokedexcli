package main

import "github.com/ArashPoorazam/pokedexcli/internal/pokeapi"

// CliCommand -
type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Config -
type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}
