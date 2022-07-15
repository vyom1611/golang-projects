package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	title string
	body string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", ":4040")
	if err != nil {
		log.Fatal("Error registering API", err)
	}

	a := Item{"First", "First Item"}
	b := Item{"Second", "Second Item"}
	c := Item{"Third", "Third Item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"Second","Second Item New"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first item: ", reply)

}