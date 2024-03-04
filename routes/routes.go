package routes

import (
	"github.com/gin-gonic/gin"
)

// Registers the routes for the application
func RegisterRoutes(server *gin.Engine) {
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.POST("/signup", signup)
}