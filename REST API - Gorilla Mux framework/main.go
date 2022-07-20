package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

// Articles endpoint
func allArticles(w http.ResponseWriter, r *http.Request) {
	//Defining articles
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test desc", Content: "Test content"},
	}
	fmt.Println("All endpoints hit")
	//Adding all articles to json write operation
	json.NewEncoder(w).Encode(articles)
}

//Trying out methods with Postman
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST endpoint reached")
}

//Homepage endpoint
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

// All requests defined here
func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", Homepage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
