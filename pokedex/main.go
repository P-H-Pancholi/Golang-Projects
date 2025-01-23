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
	next string
	prev string
	c    pokecache.Cache
}

type commandcli struct {
	name        string
	description string
	callbackfn  func(c *Config) error
}

var commandMap map[string]commandcli
var config Config

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandMap = make(map[string]commandcli)
	c := pokecache.NewCache(5 * time.Second)
	config = Config{
		next: "https://pokeapi.co/api/v2/location-area/",
		prev: "",
		c:    c,
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
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		// fmt.Printf("Your command was: %s\n", CleanInput(text)[0])
		command := CleanInput(text)[0]
		v, ok := commandMap[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			v.callbackfn(&config)
		}
	}
}

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for key, value := range commandMap {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
func GetMap(c *Config) error {
	next, prev := pokemap.GetLocArea(c.next, c.c)
	c.next = next
	c.prev = prev
	return nil
}

func GetMapb(c *Config) error {
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
