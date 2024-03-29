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
	Paid      bool                `json:"paid" bson:"paid"`
}

// Struct that represents a product in an order.
type ProductInOrder struct {
	Product string `json:"product" bson:"product" binding:"required"`
	Quantity int `json:"quantity" bson:"quantity" binding:"required"`
}

// Struct that represents a change in the paid status of an order.
type PaidStatus struct {
	Paid bool `json:"paid" bson:"paid" binding:"required"`
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
func CreateOneOrder(order *Order, loggedInUser string, ordersCollection *mongo.Collection, 
usersCollection *mongo.Collection, productsCollection *mongo.Collection) error {
	ctx := context.TODO()
	order.Paid = false // set the paid field to false automatically

	//get the user id from the email
	userInfo, err := db.GetUserByEmail(usersCollection, loggedInUser)
	if err != nil { return err }
	if userInfo == nil { return errors.New("user not found") }

	order.UserID = userInfo["_id"].(primitive.ObjectID) // set the user id in the order object automatically

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

// Updates an order in the database.
func ChangePaidStatus(id string, paidStatus *PaidStatus , collection *mongo.Collection) error {
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil { return err }

	orderIsPaid, err := OrderIsPaid(ctx, collection, objID, paidStatus.Paid)
	if err != nil { return err }
	if orderIsPaid { return errors.New("order is already paid") }

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": bson.M{"paid": true}})
	if err != nil { return err }

	return nil
}

// Checks if the order is already set to the requested paid status and returns an error if it is.
func OrderIsPaid(ctx context.Context, collection *mongo.Collection, objID primitive.ObjectID, paidStatus bool) (bool, error) {
	order := collection.FindOne(ctx, bson.M{"_id": objID})
	if order.Err() != nil { return false, order.Err() }
	
	var orderObj Order
	err := order.Decode(&orderObj)
	if err != nil { return false, err }
	
	if orderObj.Paid == paidStatus { return true, errors.New("order is already paid") }

	return false, nil
}

// Returns an order from the database if it exists, searching by id.
func GetOrderById(collection *mongo.Collection, id string) (*Order, error) {
    ctx := context.TODO()

    var order Order
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
    if err != nil {
        return nil, err
    }

    return &order, nil
}

// Updates the stock of a product in the database.
func UpdateStock(productID string, quantity int, collection *mongo.Collection) error {
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(productID)
	if err != nil { return err }

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$inc": bson.M{"stock": -quantity}})
	if err != nil { return err }

	return nil
}