package models

import (
	"context"
	"errors"
	"github.com/thegera4/cool-morning-lights-api/db"
	"github.com/thegera4/cool-morning-lights-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product struct to represent a product in the application.
type Product struct {
	ID          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name" binding:"required"`
	Description string  `bson:"description" binding:"required"`
	Price       float64 `bson:"price" binding:"required"`
	Stock       int     `bson:"stock" binding:"required"`
	Store      	string 	`bson:"store" binding:"required"`
	Pictures    []string `bson:"pictures"`
	Categories  []string `bson:"categories"`
	Active	  	bool    `bson:"active"`
}

// Struct to represent the data of a product that is going to be updated. All fields are optional.
type ProductUpdate struct {
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Stock       int     `bson:"stock"`
	Store      	string 	`bson:"store"`
	Pictures    []string `bson:"pictures"`
	Categories  []string `bson:"categories"`
	Active	  	bool    `bson:"active"`
}

// Struct to represent the data of a product that is going to be rented. All fields are required.
type RentedProduct struct {
	ID       string `bson:"id" binding:"required"`
	Quantity int    `bson:"quantity" binding:"required"`
}

// Struct to represent a slice of products that are going to be rented.
type RentRequest struct {
	RentedProducts []RentedProduct `json:"rentedProducts" binding:"required"`
}

// Returns the collection of products from the database.
func GetAllProducts() ([]Product, error) {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	var products []Product
	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil { return products, err }

	err = cursor.All(ctx, &products)
	if err != nil { return products, err }

	return products, nil
}

// Creates a product in the database.
func CreateOneProduct(product *Product) error {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	existingProduct, err := utils.ProductExistsInDb(ctx, collection, product.Name, product.Store)
	if err != nil { return err }
	if existingProduct { return errors.New("Product already exists") }

	_, err = collection.InsertOne(ctx, product)
	if err != nil { return err }

	return nil
}

// Deletes a product from the database.
func DeleteOneProduct(id string) error {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	// Convert string Id to ObjectId for MongoDB
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil { return err }

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil { return err }

	return nil
}

// Updates a product in the database.
func UpdateOneProduct(id string, product map[string]interface{}) error {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil { return err }

	// Check if there are fields to update
	if len(product) == 0 { return errors.New("no fields to update") }

	// Change the type of the stock field to int
	if stock, ok := product["stock"]; ok {
		product["stock"] = int(stock.(float64))
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": product})
	if err != nil { return err }

	return nil
}

// Updates the stock of a product in the database when an order is created.
func UpdateStockWithRent(rentedProducts []RentedProduct) error {
	collection := db.GetDBCollection("products")
	ctx := context.TODO()

	// Check if the rentedProducts slice is empty
	if len(rentedProducts) == 0 { return errors.New("no products to rent") }

	for _, rentedProduct := range rentedProducts {
		//Check if the id and quantity fields are empty
		if rentedProduct.ID == "" || rentedProduct.Quantity == 0 { return errors.New("invalid data to rent a product") }

		// Convert the string id to ObjectId
		objectID, err := primitive.ObjectIDFromHex(rentedProduct.ID)
		if err != nil { return err }

		// Get the product from the database
		var product Product
		err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
		if err != nil { return err }

		// Check if the stock is enough
		if product.Stock < rentedProduct.Quantity { return errors.New("not enough stock") }

		// Update the stock
		product.Stock -= rentedProduct.Quantity
		_, err = collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"stock": product.Stock}})
		if err != nil { return err }
	}

	return nil
}