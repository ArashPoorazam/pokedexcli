package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Command Struct
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// REPL
func startRepl() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Shows how to use this program",
			callback:    commandHelp,
		},
	}

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		input := words[0]
		if command, ok := commands[input]; !ok {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		} else {
			command.callback()
		}
	}
}

// All Commands
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
	`,
	)
	return nil
}

// Helper Functions
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
