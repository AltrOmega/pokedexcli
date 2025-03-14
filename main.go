package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"pokedexcli/pokeAPI"
	"runtime"
	"strings"
	"sync"
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
	callback    func(*config, []string) error
	config      *config
}

func commandExit(cp *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cp *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	commands := get_commands()
	for _, value := range commands {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	return nil
}

func generalMap(cp *config, goto_link string) error {
	if goto_link == "" {
		fmt.Println("End of the line.")
		return nil
	}
	got, err := pokeAPI.GetResp[pokeAPI.EnumeratedResp](goto_link)
	if err != nil {
		fmt.Println(err)
		return err
	}
	results := got.Results

	if got.Next != nil {
		cp.next = *got.Next
	} else {
		cp.next = ""
	}
	if got.Previous != nil {
		cp.previous = *got.Previous
	} else {
		cp.previous = ""
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i].Name)
	}
	return nil
}

func commandMap(cp *config, args []string) error {
	return generalMap(cp, cp.next)
}

func commandMapb(cp *config, args []string) error {
	return generalMap(cp, cp.previous)
}

func commandConfig(cp *config, args []string) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments.")
	}

	commands := get_commands()

	fmt.Println("Showing config for: ", args[0])

	val, ok := commands[args[0]]
	if !ok {
		return errors.New("No such command exists")
	}

	if val.config == nil {
		fmt.Println("Config not set for given command.")
		return nil
	}

	fmt.Printf("next: %v\n", val.config.next)
	fmt.Printf("previous: %v\n", val.config.previous)
	return nil
}

func clearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func commandClear(cp *config, args []string) error {
	return clearScreen()
}

func commandExplore(cp *config, args []string) error {
	if len(args) < 1 {
		return errors.New("Not enough arguments.")
	}
	link := fmt.Sprintf("%v%v", pokeAPI.AreaEndpoint, args[0])

	locationData, err := pokeAPI.GetResp[pokeAPI.LocationData](link)
	if err != nil {
		return err
	}

	len_ := len(locationData.PokemonEncounters)
	for i := range len_ {
		name := locationData.PokemonEncounters[i].Pokemon.Name
		fmt.Println(name)
	}

	return nil
}

var (
	commandsMap      *map[string]cliCommand
	commandsAreSetup sync.Once
)

func get_commands() map[string]cliCommand {
	commandsAreSetup.Do(func() {
		mapConfigPointer := &config{
			next:     pokeAPI.AreaEndpoint,
			previous: "",
		}
		commandsMap = &map[string]cliCommand{
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
				config:      mapConfigPointer,
			},
			"mapb": {
				name:        "mapb",
				description: "Same as map but goes in reverse",
				callback:    commandMapb,
				config:      mapConfigPointer,
			},
			"explore": {
				name:        "explore",
				description: "Lets you explore given area of the map",
				callback:    commandExplore,
			},
			"config": {
				name:        "config",
				description: "Shows configuration for a given command",
				callback:    commandConfig,
			},
			"clear": {
				name:        "clear",
				description: "Clears the termial",
				callback:    commandClear,
			},
		}
	})

	return *commandsMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := get_commands()

	for {
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

		err := command.callback(command.config, input[1:])
		commands[input[0]] = command

		if err != nil {
			fmt.Printf("Error from given command: %v\n", err)
		}
	}
}
