package main

import (
	"bufio"
	"fmt"
	"os"
)

type clicommand struct {
	name        string
	description string
	callback    func() error
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println()
		command := CleanInput(input)[0]

		if _, ok := commands[command]; !ok {
			fmt.Println("Command not found")
			continue
		}

		commands[command].callback()
	}
}

func getCommands() map[string]clicommand {

	return map[string]clicommand{
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "list commands",
			callback:    commandHelp,
		},
	}
}
