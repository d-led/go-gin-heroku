package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

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

	// https://devcenter.heroku.com/changelog-items/630
	r.GET("/version", func(c *gin.Context) {
		c.String(200, os.Getenv("SOURCE_VERSION"))
	})

	return r
}
