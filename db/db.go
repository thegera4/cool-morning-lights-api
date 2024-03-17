package db

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var mongoURI = os.Getenv("MONGO_URI")
var	dbName = os.Getenv("DB_NAME")

// Initialize the MongoDB client and return it.
func InitDB() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	options := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	var err error
	client, err = mongo.Connect(context.TODO(), options)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Connecting to MongoDB... You successfully connected to MongoDB!")

	return client
}

// Returns a collection from the Mongo database or create and return a collection if it doesn't exist.
func GetDBCollection(collectionName string) *mongo.Collection {
	if client == nil {
		panic("MongoDB client is not initialized. Call InitDB first.")
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	names, err := database.ListCollectionNames(context.Background(), bson.M{"name": collectionName})
	if err != nil {
		panic(err)
	}
	// If the collection is not in the list, create it
	if len(names) == 0 {
		createOpts := options.CreateCollection().SetValidator(bson.M{})
		if err := database.CreateCollection(context.Background(), collectionName, createOpts); err != nil {
			panic(err)
		}
	}

	return collection
}

// Return a user from the database if it exists, searching by email.
func GetUserByEmail(collection *mongo.Collection, email string) (bson.M, error) {
	var user bson.M
	err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Return a user from the database if it exists, searching by id.
func GetUserById(collection *mongo.Collection, id string) (bson.M, error) {
	var user bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Returns all users that have the "admin" field set to true.
func GetAdminUsers(collection *mongo.Collection) ([]bson.M, error) {
	var users []bson.M
	cursor, err := collection.Find(context.TODO(), bson.M{"admin": true})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Returns a product from the database if it exists, searching by id.
func GetProductById(collection *mongo.Collection, id string) (bson.M, error) {
	var product bson.M
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}