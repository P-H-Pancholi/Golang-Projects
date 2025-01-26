package pokemap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/P-H-Pancholi/Golang-Projects/pokedex/pokecache"
)

func GetLocArea(url string, c pokecache.Cache) (next string, prev string) {
	var body []byte
	body, ok := c.Get(url)
	if !ok {
		body := ApiCall(url)
		c.Add(url, body)
	}

	locArea := LocAreaStruct{}
	if err := json.Unmarshal(body, &locArea); err != nil {
		log.Fatal(err)
	}
	for _, r := range locArea.Results {
		fmt.Println(r.Name)
	}
	return locArea.Next, locArea.Previous
}

func ExploreArea(location string, c pokecache.Cache) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	var body []byte
	body, ok := c.Get(url)
	if !ok {
		body := ApiCall(url)
		c.Add(url, body)
	}

	poke := PokemonAreaEncounters{}

	if err := json.Unmarshal(body, &poke); err != nil {
		log.Fatal(err)
	}
	for _, r := range poke.PokemonEncounters {
		fmt.Println(r.Pokemon.Name)
	}

}

func CatchPoke(text string, m map[string]Pokemon) bool {
	url := "https://pokeapi.co/api/v2/pokemon/" + text

	body := ApiCall(url)
	var p Pokemon
	if err := json.Unmarshal(body, &p); err != nil {
		log.Fatalf("unable to unmarshal response body : %s", err)
	}

	n := rand.Intn(700)
	if p.BaseExperience < n {
		m[text] = p
		return true
	} else {
		return false
	}
}

func ApiCall(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("unable to call endpoint : %s", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("unable to read response body : %s", err)
	}
	defer res.Body.Close()
	return body
}
