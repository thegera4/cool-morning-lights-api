package db

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
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
	// Check if the client is initialized
	if client == nil {
		panic("MongoDB client is not initialized. Call InitDB first.")
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	// List all collections in the database
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