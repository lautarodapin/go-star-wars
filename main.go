package main

import (
	"star-wars-api/people"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/people", people.PeopleList())
	r.GET("/people/:id", people.PeopleDetail())
	r.Run() // listen and serve on 0.0.0.0:8080
}
