package main

import "github.com/gin-gonic/gin"

func main() {
	server().Run()
}

func server() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})
	return r
}
