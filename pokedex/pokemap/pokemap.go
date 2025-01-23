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
