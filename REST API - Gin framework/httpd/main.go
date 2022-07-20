package main

import (
	"github.com/gin-gonic/gin"
	"main.go/httpd/handler"
)

func main() {
	//Making the router through Gin
	r := gin.Default()

	//Defining HTTP methods using created handler functions
	r.GET("/ping", handler.PingGet())
	r.POST("/newsfeed", handler.NewsfeedPost())
	r.GET("/newsfeed", handler.NewsfeedGet())

	//Running the server
	r.Run()

}
