package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	s := strings.ToLower(text)
	return strings.Fields(s)
}

type config struct {
	next     string
	previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(cp *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cp *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, value := range get_commands() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}

func commandMap(cp *config) error {

}

func get_commands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays name of 20 locations in the Pokemon world. Each subsequent call displays the next 20 locations.",
			callback:    commandMap,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := get_commands()

	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) < 1 {
			fmt.Println("Unknown command")
			continue
		}

		command, ok := commands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Errorf("Error from given command: %w\n", err)
		}
	}
}
