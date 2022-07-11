package main

import (
	"fmt"
	"log"
	"net/http"
)

//
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Printf("Name: %v, Address: %v", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	//Assigning path for static page
	var fileServer = http.FileServer(http.Dir("./static"))

	//Setting routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server started on port 8080\n")
	var err any
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
