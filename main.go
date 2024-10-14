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
		"exit": {
			name:			"exit",
			description:	"Exit the pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name: 			"map",
			description:	"Displays the next 20 map locations",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays the previous 20 map locations",
			callback:		commandMapb,
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
	os.Exit(0)
	return nil
}

func commandMap() (error) {
	return nil
}

func commandMapb() (error) {
	return nil
}

func main() {
	commands := getCommands()
	for {
		fmt.Print("pokedex > ")
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
