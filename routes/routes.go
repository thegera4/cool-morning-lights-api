package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/middlewares"
)

// Registers the routes for the application.
func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.PATCH("/users/validate", validate)

	// Admin Protected routes
	admin := server.Group("/")
	admin.Use(middlewares.AdminRequired)
	admin.GET("/users", getUsers)
	admin.PATCH("/users/makeAdmin", makeAdmin)

	// User Protected routes
	user := server.Group("/")
	user.Use(middlewares.Authenticate)
	user.PATCH("/users/updateInfo", updateInfo)
}