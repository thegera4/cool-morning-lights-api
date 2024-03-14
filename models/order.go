package models

import (
	"context"
	"github.com/thegera4/cool-morning-lights-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct that represents a purchase/rent order.
type Order struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Products  []primitive.ObjectID `json:"products" bson:"products"`
	Total     float64            `json:"total" bson:"total"`
	Store     primitive.ObjectID `json:"store" bson:"store"`
	Start	  string             `json:"start" bson:"start"`
	End		  string             `json:"end" bson:"end"`
}

// Returns the collection of orders from the database.
func GetAllOrders() ([]Order, error) {
	collection := db.GetDBCollection("orders")
	ctx := context.TODO()

	var orders []Order
	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil { return orders, err }

	err = cursor.All(ctx, &orders)
	if err != nil { return orders, err }

	return orders, nil
}

// Deletes an order from the database.
func DeleteOneOrder(id string) error {
	collection := db.GetDBCollection("orders")
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil { return err }

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil { return err }

	return nil
}