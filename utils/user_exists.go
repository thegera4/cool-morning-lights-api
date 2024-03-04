package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Returns true if the user exists in the Mongo database.
func UserExistsInDb(ctx context.Context, collection *mongo.Collection, email string) (bool, error) {
	var existingUser bson.M

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}