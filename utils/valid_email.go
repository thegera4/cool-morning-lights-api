package utils

import (
	"regexp"
)

// Checks if the given email is valid.
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-z]{2,4}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}