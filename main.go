package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() (map[string]cliCommand) {
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
		},
	}
}

func commandHelp() (error) {
	fmt.Println("commandHelp")
	return nil
}

func commandExit() (error){
	fmt.Println("commandExit")
	return nil
}

func main() {
	//commands := getCommands()
	for {
		// get input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println("In scan")
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("reading standard input:", err)
		}
	}
}
