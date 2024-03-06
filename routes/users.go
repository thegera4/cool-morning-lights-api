package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/models"
	"github.com/thegera4/cool-morning-lights-api/utils"
)

/* Request Handlers */

// Handles the signup (register) request for new users.
func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err := user.Save(); 
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save user"})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Handles the login request for users. Returns a JWT token if the credentials are correct.
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user);
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Logged in successfully!", "token": token})
}

// Handles the request to validate an account.
func validate(context *gin.Context) {
	type validateRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	var req validateRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!" + err.Error()})
		return
	}

	err := models.ValidateAccount(req.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to validate account"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Account validated successfully"})
}

// Handles the request to get all users.
func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}