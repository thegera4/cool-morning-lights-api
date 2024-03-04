package models

import (
	"context"
	"errors"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Username string
}

// Saves a user to the database if it doesn't exist
func (u User) Save() error {

	// Get or Create a Mongo collection for users
	collection := db.GetDBCollection("users")
	// Create a context
	ctx := context.TODO()

	// Validate if the user already exists
	var existingUser User
	err := collection.FindOne(ctx, bson.M{"email": u.Email}).Decode(&existingUser)
	if err == nil {
		return errors.New("User already exists")
	} else if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Insert the user into the collection
	_, err = collection.InsertOne(ctx, bson.M{
		"email":    u.Email,
		"password": hashedPassword,
		"username": u.Username,
	})
	if err != nil {
		return err
	}

	return nil
}