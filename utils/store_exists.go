package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Returns true if the store exists in the "stores" collection in the MongoDB.
func StoreExistsInDb(ctx context.Context, collection *mongo.Collection, name string) (bool, error) {
	var existingStore bson.M

	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&existingStore)
	if err != nil {
		if err == mongo.ErrNoDocuments { return false, nil }
		return false, err
	}

	return true, nil
}