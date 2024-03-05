package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Returns true if the username exists in the Mongo database.
func UsernameExistsInDb(ctx context.Context, collection *mongo.Collection, username string) (bool, error) {
	var existingUser bson.M

	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}