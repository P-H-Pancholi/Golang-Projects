package pokemap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

func GetLocArea(url string) (next string, prev string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
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
