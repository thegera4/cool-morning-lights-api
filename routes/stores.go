package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
	"github.com/thegera4/cool-morning-lights-api/utils"
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

// Handles the request to create a store.
func createStore(c *gin.Context) {
	var store models.Store
	err := c.BindJSON(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	store.Active = true

	err = models.CreateOneStore(&store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Store created successfully"})
}

// Handles the request to update a store.
func updateStore(c *gin.Context) {
	id := c.Param("id")
	var store map[string]interface{}

	err := c.BindJSON(&store)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	} 

	validUpdate := utils.StoreUpdateIsValid(store)
	if !validUpdate { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = models.UpdateOneStore(id, store)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store updated successfully"})
}
