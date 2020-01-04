package main

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	// Return message "Hello World!" in JSON format
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func HelloName(c *gin.Context) {
	// Passing parameter :name  to output message "Hello :name!" in JSON format
	name := c.Param("name")
	message := "Hello " + name + "!"
	c.JSON(200, gin.H{
		"message": message,
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", Hello)
	r.GET("/greet/:name", HelloName)
	r.Run()
}
