package main

import (
	"github.com/gin-gonic/gin"
	"main.go/httpd/handler"
	"main.go/platform/newsfeed"
)

func main() {
	//Making the router through Gin
	r := gin.Default()

	//Generating a new feed item defined by the JSON schema for handler function params
	feed := newsfeed.New()

	//Defining HTTP methods using created handler functions
	r.GET("/ping", handler.PingGet())
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.GET("/newsfeed", handler.NewsfeedGet(feed))

	//Running the server
	r.Run()

}
