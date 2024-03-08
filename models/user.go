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
	Id       string `bson:"_id,omitempty"` // The ID field is treated as an ObjectId to work with MongoDB
	Email    string `bson:"email" binding:"required"`
	Password string `bson:"password" binding:"required"`
	Username string `bson:"username"`
	Admin  	  bool  `bson:"admin"`
	Validated bool 	`bson:"validated"`
}

// Saves a user to the Mongo database if it doesn't exist.
func (u User) Save() error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	// Use channels to communicate the results of each validation
	emailValidated := make(chan bool)
	usernameValidated := make(chan bool)
	userValidated := make(chan bool)

	// Concurrently validate the email, username and password
	go func() { emailValidated <- utils.IsValidEmail(u.Email) }()

	go func() {
		exists, err := utils.UsernameExistsInDb(ctx, collection, u.Username)
		if err != nil { 
			usernameValidated <- false 
			return
		}
		usernameValidated <- exists
	}()

	go func() {
		exists, err := utils.UserExistsInDb(ctx, collection, u.Email)
		if err != nil {
			userValidated <- false
			return
		}
		userValidated <- exists
	}()

	// Wait for the results of the validations and close the channels
	emailIsValid := <-emailValidated
	close(emailValidated)
	usernameExists := <-usernameValidated
	close(usernameValidated)
	userExists := <-userValidated
	close(userValidated)

	// Check validation results
	if !emailIsValid { return errors.New("invalid email") }
	if userExists { return errors.New("user already exists") }
	if usernameExists { return errors.New("username already exists") }

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil && err.Error() == "password must be at least 8 characters long" {
		return errors.New("password must be at least 8 characters long")
	}
	if err != nil { return err }

	_, err = collection.InsertOne(ctx, bson.M{
		"email":    u.Email,
		"password": hashedPassword,
		"username": u.Username,
		"admin":    false,
		"validated": false,
	})
	if err != nil { return err }

	return nil
}

// Updates the user's information by changing the "username" and/or "password" fields.
func (u User) UpdateUserInfo() error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	user, err := db.GetUserByEmail(collection, u.Email)
	if err != nil { return err }

	// Validations to check if the username and password are the same as the current ones, and theres at least one field to update
	if u.Username == "" && u.Password == "" { return errors.New("no fields to update") }
	if user["username"].(string) == u.Username { return errors.New("username is the same as the current username") }
	if utils.CheckPasswordHash(u.Password, user["password"].(string)) { return errors.New("password is the same as the current password") }

	// If everything is valid, and if we have a username and or password, update the user's information
	if u.Username != "" {
		_, err = collection.UpdateOne(ctx, bson.M{"email": u.Email}, bson.M{"$set": bson.M{"username": u.Username}})
		if err != nil { return err }
	}

	if u.Password != "" {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil && err.Error() == "password must be at least 8 characters long" {
			return errors.New("password must be at least 8 characters long")
		}
		if err != nil { return err }
		_, err = collection.UpdateOne(ctx, bson.M{"email": u.Email}, bson.M{"$set": bson.M{"password": hashedPassword}})
		if err != nil { return err }
	}

	return nil
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

// User struct without the Password field, to be used when returning users from the database.
type UserWithoutPassword struct {
    Email    string `bson:"email"`
    Username string `bson:"username"`
	Admin    bool   `bson:"admin"`
	Validated bool  `bson:"validated"`
}

// Returns all users from the Mongo database without including their passwords.
func GetAllUsers() ([]UserWithoutPassword, error) {
	var users []UserWithoutPassword
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	projection := bson.M{"email": 1, "username": 1, "admin": 1, "validated": 1} // Exclude the password from the query
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

// Validates the user's account by changing the "validated" field to true.
func ValidateAccount(email string) error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	userExists, err := utils.UserExistsInDb(ctx, collection, email)
	if err != nil { return err }
	if !userExists { return errors.New("User does not exist") }

	_, err = collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"validated": true}})
	if err != nil { return err }

	return nil
}

// Changes the user's role to admin by updating the "admin" field to true.
func MakeUserAdmin(email string) error {
	collection := db.GetDBCollection("users")
	ctx := context.TODO()

	userExists, err := utils.UserExistsInDb(ctx, collection, email)
	if err != nil { return err }
	if !userExists { return errors.New("User does not exist") }

	_, err = collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"admin": true}})
	if err != nil { return err }

	return nil
}