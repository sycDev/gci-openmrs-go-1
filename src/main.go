package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func Greet(c *gin.Context) {
	name := c.Param("name")
	message := "Hello " + name + "!"
	c.String(http.StatusOK, message)
}

func main() {	
	r := gin.Default()
	r.GET("/hello", Hello)
	r.GET("/greet/:name", Greet)
	r.Run()
}
