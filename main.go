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
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands := getCommands()
	for k := range commands {
		fmt.Print(k)
		fmt.Print(": ")
		fmt.Println(commands[k].description)
	}

	return nil
}

func commandExit() (error){			
	fmt.Println("")

	os.Exit(0)
	return nil
}

func main() {
	commands := getCommands()
	for {
		fmt.Print("pokedex > ")
		// get input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Print("\npokedex > ")
			for k := range commands {
				if scanner.Text() == k {
					commands[k].callback()
					fmt.Print("\npokedex > ")
				}
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("reading standard input:", err)
		}
	}
}
