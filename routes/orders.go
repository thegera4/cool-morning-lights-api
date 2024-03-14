package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
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