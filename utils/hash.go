package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Hashes (encrypts) a password for safe storage in the database.
func HashPassword(password string) (string, error) {
	validPassword := len(password) >= 8
	if !validPassword {
		return "Password must be at least 8 characters long", errors.New("password must be at least 8 characters long")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Compares a password with its hash to confirm if they match.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}