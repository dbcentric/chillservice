package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	server := gin.Default()

	server.GET("api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": uuid.New().String(),
		})
	})

	server.Run()
}
