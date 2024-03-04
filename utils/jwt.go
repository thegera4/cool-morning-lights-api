package utils

import (
	"errors"
	"os"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

// Generates a new JWT token.
func GenerateToken(email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// Validates a JWT token and returns the user's id.
func ValidateToken(token string) (int64, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok { 
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("the token fails here")
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64)) //convert it

	return userId, nil
}