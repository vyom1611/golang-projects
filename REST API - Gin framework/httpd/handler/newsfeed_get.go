package handler

import (
	"github.com/gin-gonic/gin"
	"main.go/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Getting all items
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
