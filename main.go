package main

import (
	"time"
)

func main() {
	pokeClient := NewClient(20 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
