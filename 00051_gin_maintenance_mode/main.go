package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var maintenanceMode = false

type updateMaintenance struct {
	Value bool
}

func main() {
	g := gin.Default()

	public := g.Group("/")
	{
		public.Use(maintenanceMiddleware)
		public.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World!",
			})
		})
	}

	internal := g.Group("/internal")
	{
		internal.GET("/maintenance", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name":  "maintenanceMode",
				"value": maintenanceMode,
			})
		})
		internal.POST("/maintenance", func(c *gin.Context) {
			var body updateMaintenance
			if err := c.BindJSON(&body); err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			maintenanceMode = body.Value
			c.AbortWithStatus(http.StatusNoContent)
		})
	}

	_ = g.Run("localhost:8080")
}

func maintenanceMiddleware(c *gin.Context) {
	if maintenanceMode {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"message": "currently in maintenance mode",
		})
		return
	}

	c.Next()
}
