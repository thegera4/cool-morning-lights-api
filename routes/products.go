package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
)

/* Request Handlers */

// Handles the request to get all products.
func getProducts(context *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	context.JSON(http.StatusOK, products)
}

// Handles the request to create a product.
func createProduct(context *gin.Context) {
	var product models.Product
	err := context.BindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = models.CreateOneProduct(&product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	context.JSON(http.StatusCreated, product)
}

// Handles the request to delete a product.
func deleteProduct(context *gin.Context) {
	id := context.Param("id")
	err := models.DeleteOneProduct(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// Handles the request to update a product.
func updateProduct(context *gin.Context) {
	id := context.Param("id")
	var product models.Product
	err := context.BindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err = models.UpdateOneProduct(id, &product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	context.JSON(http.StatusOK, product)
}