package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/P-H-Pancholi/Golang-Projects/pokedex/pokecache"
	"github.com/P-H-Pancholi/Golang-Projects/pokedex/pokemap"
)

type Config struct {
	next          string
	prev          string
	c             pokecache.Cache
	pokemonCaught map[string]pokemap.Pokemon
}

type commandcli struct {
	name        string
	description string
	callbackfn  func(c *Config, s string) error
}

var commandMap map[string]commandcli
var config Config

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap = make(map[string]commandcli)
	c := pokecache.NewCache(5 * time.Second)
	config = Config{
		next:          "https://pokeapi.co/api/v2/location-area/",
		prev:          "",
		c:             c,
		pokemonCaught: make(map[string]pokemap.Pokemon),
	}
	commandMap["exit"] = commandcli{
		name:        "exit",
		description: "Exit the Pokedex",
		callbackfn:  commandExit,
	}
	commandMap["help"] = commandcli{
		name:        "help",
		description: "Displays a help message",
		callbackfn:  commandHelp,
	}
	commandMap["map"] = commandcli{
		name:        "map",
		description: "displays the names of next 20 location areas in the Pokemon world",
		callbackfn:  GetMap,
	}
	commandMap["mapb"] = commandcli{
		name:        "mapb",
		description: "displays the names of prev 20 location areas in the Pokemon world",
		callbackfn:  GetMapb,
	}
	commandMap["explore"] = commandcli{
		name:        "explore",
		description: "Explores pokemon in area provided",
		callbackfn:  Explore,
	}
	commandMap["catch"] = commandcli{
		name:        "catch",
		description: "catch specified pokemon ",
		callbackfn:  Catch,
	}
	commandMap["inspect"] = commandcli{
		name:        "inspect",
		description: "It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
		callbackfn:  Inspect,
	}
	commandMap["pokedex"] = commandcli{
		name:        "pokedex",
		description: " It takes no arguments but prints a list of all the names of the Pokemon the user has caught.",
		callbackfn:  PrintPokemon,
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		// fmt.Printf("Your command was: %s\n", CleanInput(text)[0])
		command := CleanInput(text)
		v, ok := commandMap[command[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if len(command) > 1 {
				v.callbackfn(&config, command[1])
			} else {
				v.callbackfn(&config, "")
			}

		}
	}
}

func commandExit(c *Config, s string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *Config, s string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, value := range commandMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
func GetMap(c *Config, s string) error {
	next, prev := pokemap.GetLocArea(c.next, c.c)
	c.next = next
	c.prev = prev
	return nil
}

func GetMapb(c *Config, s string) error {
	if c.prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	next, prev := pokemap.GetLocArea(c.prev, c.c)
	c.next = next
	c.prev = prev
	return nil
}

func CleanInput(text string) []string {
	text = strings.ToLower(text)

	return strings.Fields(text)
}

func Explore(c *Config, s string) error {
	pokemap.ExploreArea(s, c.c)
	return nil
}

func Catch(c *Config, s string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", s)
	if pokemap.CatchPoke(s, c.pokemonCaught) {
		fmt.Printf("%s was caught!\n", s)
	} else {
		fmt.Printf("%s escaped!\n", s)
	}
	return nil
}

func Inspect(c *Config, s string) error {
	v, ok := c.pokemonCaught[s]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats: \n", v.Name, v.Height, v.Weight)
		for _, val := range v.Stats {
			fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range v.Types {
			fmt.Printf("  - %s\n", val.Type.Name)
		}
	}
	return nil
}

func PrintPokemon(c *Config, s string) error {
	fmt.Println("Your Pokedex:")
	for key := range c.pokemonCaught {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
