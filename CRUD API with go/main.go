package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1",Isbn:"32423",Title:"Banana", Director: &Director{Firstname: "John", Lastname: "Williams"} })
	movies = append(movies, Movie{ID:"2",Isbn:"02847",Title:"Apple", Director: &Director{Firstname: "Bob", Lastname: "Zickovic"} })
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started at port at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
