package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/thegera4/cool-morning-lights-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		//AbortWithStatusJSON is a helper function that aborts a request with the specified status code and payload in JSON format
		//no other handlers will be called from this point
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	email, err := utils.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	context.Set("email", email) //set the email in the context (add data to the context)
	context.Next() //call the next handler
}