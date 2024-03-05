package models

import (
	"context"
	"errors"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct to represent a user in the application
type User struct {
	//Id       int64 no need to include this field, MongoDB will automatically generate an ID for each user
	Email    string `bson:"email" binding:"required"`
	Password string `bson:"password" binding:"required"`
	Username string `bson:"username"`
}

// Saves a user to the Mongo database if it doesn't exist.
func (u User) Save() error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	userExists, err := utils.UserExistsInDb(ctx, collection, u.Email)
	if err != nil { return err }
	if userExists { return errors.New("User already exists") }

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil && err.Error() == "password must be at least 8 characters long" {
		return errors.New("password must be at least 8 characters long")
	}
	if err != nil { return err }

	_, err = collection.InsertOne(ctx, bson.M{
		"email":    u.Email,
		"password": hashedPassword,
		"username": u.Username,
	})
	if err != nil { return err }

	return nil
}

// User struct without the Password field, to be used when returning users from the database.
type UserWithoutPassword struct {
    Email    string `bson:"email"`
    Username string `bson:"username"`
}

// Returns all users from the Mongo database without including their passwords.
func GetAllUsers() ([]UserWithoutPassword, error) {
	var users []UserWithoutPassword
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	projection := bson.M{"email": 1, "username": 1} // Exclude the password from the query
	cursor, err := collection.Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil { return nil, err }

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserWithoutPassword
		err := cursor.Decode(&user)
		if err != nil { return nil, err }
		users = append(users, user)
	}

	return users, nil
}

// Validates the user's credentials by checking if the user exists and if the password is correct.
func (u User) ValidateCredentials() error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	userExists, err := utils.UserExistsInDb(ctx, collection, u.Email)
	if err != nil { return err }
	if !userExists { return errors.New("User does not exist") }

	userData, err := db.GetUserByEmail(collection, u.Email)
	if err != nil { return err }

	passwordIsValid := utils.CheckPasswordHash(u.Password, userData["password"].(string))
	if !passwordIsValid { return errors.New("invalid credentials") }

	return nil
}