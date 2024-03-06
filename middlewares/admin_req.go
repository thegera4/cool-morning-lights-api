package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/utils"
)

// Middleware to confirm if the user is an admin.
func AdminRequired(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

    if token == "" {
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
        return
    }

    email, err := utils.ValidateToken(token)
    if err != nil {
        context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
        return
    }

	collection := db.GetDBCollection("users")
	if collection == nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to connect to database!"})
		return
	}

	adminUsers, err := db.GetAdminUsers(collection)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to get admin users!"})
		return
	}

	isAdmin := false
	for _, adminUser := range adminUsers {
		if adminUser["email"].(string) == email {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

    context.Set("email", email)
    context.Next()
}