package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/models"
)

/* Request Handlers */

// Handles the request to get all orders.
func getOrders(c *gin.Context) {
	orders, err := models.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// Handles the request to delete an order.
func deleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteOneOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// Handles the request to create an order.
func createOrder(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request" + err.Error()})
		return
	}

	loggedInUser := c.GetString("email")

	// TODO:concurrency
	usersCollection := db.GetDBCollection("users")
	productsCollection := db.GetDBCollection("products")
	ordersCollection := db.GetDBCollection("orders")

	err = models.CreateOneOrder(&order, loggedInUser, ordersCollection, usersCollection, productsCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}