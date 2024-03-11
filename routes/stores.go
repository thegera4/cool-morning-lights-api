package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
)

/* Request Handlers */

// Handles the request to get all stores.
func getStores(c *gin.Context) {
	stores, err := models.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stores)
}

// Handles the request to delete a store.
func deleteStore(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteOneStore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete store"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}
