package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Returns true if the product exists in the "products" collection in the MongoDB.
func ProductExistsInDb(ctx context.Context, collection *mongo.Collection, name string, store string) (bool, error) {
	var existingProduct bson.M

	err := collection.FindOne(ctx, bson.M{"name": name, "store": store}).Decode(&existingProduct)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}