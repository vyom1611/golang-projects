package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)


type Item struct {
	title string
	body string
}

var database []Item

func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

type API int

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}


func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)

	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}

	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	
	*reply = del
	return nil
}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp",":4040")
	if err != nil {
		log.Fatal("Error serving: ", err)
	}

	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving:", err)
	}

}

