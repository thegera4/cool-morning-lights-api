package routes

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
	//"github.com/thegera4/cool-morning-lights-api/utils"
)

/* Request Handlers */

// Handles the signup (register) request for new users.
func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil { // Bind the request body to the user model
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := user.Save(); err != nil { // Save the user to the database
		context.JSON(500, gin.H{"error": "Failed to save user"})
		return
	}

	context.JSON(201, gin.H{"message": "User created successfully"})
}

// Handles the request to get all users.
func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers() // Get all users from the database
	if err != nil {
		context.JSON(500, gin.H{"error": "Failed to get users"})
		return
	}

	context.JSON(200, gin.H{"users": users})
}