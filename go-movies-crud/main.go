package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

// Handler func are of type fun (ResponseWriter, *Request)
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set header values for ResponseWriter
	w.Header().Set("Content-Type", "application/json")

	//Writes movies with JSON encoding to ResponseWriter
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// will return parameters from the url path

	for key, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:key], movies[key+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// will return parameters from the url path

	// Range over movies slice to find movie through id
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	//json decoding the request body & writing values to movie var
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for key, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:key], movies[key+1:]...)
			break
		}
	}

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}
func main() {
	r := mux.NewRouter()
	// mux is a HTTP request multiplexer, used for routing requests

	movies = append(movies, Movie{ID: "1", Isbn: "34523477643", Title: "Movie One", Director: &Director{FirstName: "ken", LastName: "Adams"}})
	movies = append(movies, Movie{ID: "2", Isbn: "56348596857", Title: "Movie Two", Director: &Director{FirstName: "John", LastName: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	// Handler function to route GET /movies to getMovies func
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// Handler function to route GET /movies/id to getMovie func
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// Handler function to route POST /movies to createMovie func
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// Handler function to route PUT /movies/id to updateMovie func
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	// Handler function to route DELETE /movies/id to deleteMovie func

	fmt.Printf("....Starting server at port 8000\n")

	if err := http.ListenAndServe(":8000", r); err != nil {
		// Listen&Serve will start a web server on specified address
		log.Fatal(err)
	}

}
