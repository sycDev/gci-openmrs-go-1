package main

import (
	"github.com/gin-gonic/gin"
)

const Name = "name"

// Return message "Hello World!" in JSON format
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

// Return message "Hello" and name parameter in JSON format
func Greet(c *gin.Context) {
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
