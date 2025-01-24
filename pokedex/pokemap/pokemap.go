package pokemap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/P-H-Pancholi/Golang-Projects/pokedex/pokecache"
)

type LocAreaStruct struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonAreaEncounters struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocArea(url string, c pokecache.Cache) (next string, prev string) {
	var body []byte
	body, ok := c.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
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
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
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
