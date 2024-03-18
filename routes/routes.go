package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/middlewares"
)

// Registers the routes for the application.
func RegisterRoutes(server *gin.Engine) {
	// Public routes
	server.POST("/signup", Signup)
	server.POST("/login", login)
	server.PATCH("/users/validate", validate)
	server.GET("/products", getProducts)
	server.GET("/api/stores", getStores)

	// Admin Protected routes
	admin := server.Group("/")
	admin.Use(middlewares.AdminRequired)
	admin.GET("/users", getUsers)
	admin.PATCH("/users/makeAdmin", makeAdmin)
	admin.POST("/products", createProduct)
	admin.DELETE("/products/:id", deleteProduct)
	admin.PUT("/products/:id", updateProduct)
	admin.POST("/api/stores", createStore)
	admin.DELETE("/api/stores/:id", deleteStore)
	admin.PUT("/api/stores/:id", updateStore)
	admin.GET("/api/orders", getOrders)
	admin.DELETE("/api/orders/:id", deleteOrder)
	admin.PATCH("/api/orders/:id", updateOrder)

	// User Protected routes
	user := server.Group("/")
	user.Use(middlewares.Authenticate)
	user.PATCH("/users/updateInfo", updateInfo)
	user.PATCH("/products", rentProducts)
	user.POST("/api/orders", createOrder)
}