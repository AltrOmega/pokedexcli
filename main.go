package main

/*
Todo:
  clean up the code and move some thing to a new package
  add mapb
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "pokedexcli/pokeAPI"
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
  config      config
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
  fmt.Printf("runing link: %v\n", cp.next)
  got := pokeAPI.GetLocationArea(cp.next)
  results := got.Results
  //pokeAPI.EnumeratedResp

  fmt.Printf("next: %v\n", cp.next)
  fmt.Printf("previous: %v\n", cp.previous)

  if got.Next != nil {
    cp.next = *got.Next
  }
  if got.Previous != nil {
    cp.previous = *got.Previous
  }

  fmt.Printf("next: %v\n", cp.next)
  fmt.Printf("previous: %v\n", cp.previous)
  
  for i := 0; i < len(results); i++ {
    fmt.Println( results[i].Name )
  }
  return nil
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
      config:   config{
            next: "https://pokeapi.co/api/v2/location-area/",
            previous: "",
        },
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

    fmt.Println(command.config)
		err := command.callback(&command.config)

    fmt.Printf("out next: %v\n", command.config.next)
    fmt.Printf("out previous: %v\n", command.config.previous)
    commands[input[0]] = command

		if err != nil {
			fmt.Errorf("Error from given command: %w\n", err)
		}
	}
}
