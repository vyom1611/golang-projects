package main

import (
	"fmt"
	"log"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func HandleRequest() {
	http.HandleFunc("/", Homepage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	
}