package models

import (
	"context"
	"github.com/thegera4/cool-morning-lights-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product struct to represent a product in the application.
type Product struct {
	Id          string  `bson:"_id,omitempty"` // The ID field is treated as an ObjectId to work with MongoDB
	Name        string  `bson:"name" binding:"required"`
	Description string  `bson:"description" binding:"required"`
	Price       float64 `bson:"price" binding:"required"`
	Stock       int     `bson:"stock" binding:"required"`
	Pictures   []string `bson:"pictures"`
	Stores      []string `bson:"stores"`
}

// Returns the collection of products from the database.
func GetAllProducts() ([]Product, error) {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	var products []Product
	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return products, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return products, err
	}

	return products, nil
}