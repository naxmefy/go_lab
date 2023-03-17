package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	_ = g.Run("localhost:8080")
}
