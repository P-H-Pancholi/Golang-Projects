package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
	}
	fmt.Fprintf(w, "Post request Successfull\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Handler func
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not accepted", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello, Welcome to the Page!")
}

func main() {
	// Will return an Handler with contents of fs rooted at root
	fileServer := http.FileServer(http.Dir("./static"))

	//Handle & HandlerFunc adds handler to def handler (DefaultServeMux)
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000")

	//Listen&Serve starts HTTP server with a given addr & handler (def : DefaultServeMux)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
