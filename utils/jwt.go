package utils

import (
	"errors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

// Generates a new JWT token.
func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// Validates a JWT token and returns the email of the user.
func ValidateToken(token string) (string,error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok { return nil, errors.New("unexpected signing method") }
		return []byte(secretKey), nil
	})
	if err != nil { return "", errors.New("could not parse token")}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid { return "", errors.New("invalid token") }

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok { return "", errors.New("invalid token claims") }

	email := claims["email"].(string)

	return email, nil
}