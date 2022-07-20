package handler

import (
	"github.com/gin-gonic/gin"
	"main.go/platform/newsfeed"
	"net/http"
)

//Defining schema for request processing
type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Defining request body using the schema to bind to the gin context
		requestBody := newsfeedPostRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			return
		}

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
