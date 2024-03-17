package models

import (
	"context"
	"errors"
	"github.com/thegera4/cool-morning-lights-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Struct that represents a purchase/rent order.
type Order struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty"`
	UserID    primitive.ObjectID  `json:"userId" bson:"userId"`
	Products  []ProductInOrder    `json:"products" bson:"products" binding:"required"`
	Total     float64             `json:"total" bson:"total"`
	Store     string 			  `json:"store" bson:"store" binding:"required"`
	Start	  string              `json:"start" bson:"start" binding:"required"`
	End		  string              `json:"end" bson:"end" binding:"required"`
}

// Struct that represents a product in an order.
type ProductInOrder struct {
	Product string `json:"product" bson:"product" binding:"required"`
	Quantity int `json:"quantity" bson:"quantity" binding:"required"`
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

// Creates an order in the database.
func CreateOneOrder(order *Order, loggedInUser string, usersCollection *mongo.Collection, productsCollection *mongo.Collection) error {
	ordersCollection := db.GetDBCollection("orders")
	ctx := context.TODO()

	//get the user id from the email
	userInfo, err := db.GetUserByEmail(usersCollection, loggedInUser)
	if err != nil { return err }
	if userInfo == nil { return errors.New("user not found") }

	// set the user id in the order object to automatically link the order to the user
	order.UserID = userInfo["_id"].(primitive.ObjectID)

	// calculate the total of the order based on the products and their quantities
	var total float64
	for _, product := range order.Products {
		productInfo, err := db.GetProductById(productsCollection, product.Product)
		if err != nil { return err }
		if productInfo == nil { return errors.New("product not found") }

		total += productInfo["price"].(float64) * float64(product.Quantity)
	}

	// set the total in the order object
	order.Total = total

	_, err = ordersCollection.InsertOne(ctx, order)
	if err != nil { return err }

	return nil
}