package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
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


//Homepage endpoint
func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

// All requests defined here
func handleRequest() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}


func main() {
	handleRequest()
}