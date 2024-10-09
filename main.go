package main

import (
	"bufio"
	"fmt"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

return map[string]cliCommand {
	"help": {
		name:			"help",
		description:	"Displays a help message",
		callback:		commandHelp,
	},
	"exit": {
		name:			"exit",
		description:	"Exit the pokedex",
		callback:		commandExit,
	}
}

func commandHelp() {
	fmt.Println("commandHelp")
}

func commandExit() {
	fmt.Println("commandExit")
}

func main() {
	for {

	}
}
