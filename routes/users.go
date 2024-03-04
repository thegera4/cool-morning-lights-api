package routes

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
	//"github.com/thegera4/cool-morning-lights-api/utils"
)

/* request handlers */

// Handles the signup (register) request for new users
func signup(context *gin.Context) {
	var user models.User

	// Bind the request body to the user model
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Save the user to the database
	if err := user.Save(); err != nil {
		context.JSON(500, gin.H{"error": "Failed to save user"})
		return
	}

	context.JSON(201, gin.H{"message": "User created successfully"})
	
}