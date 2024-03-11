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

// Store struct to represent a store in the application.
type Store struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	Name        string  `bson:"name" binding:"required"`
	Address    	string  `bson:"location" binding:"required"`
	ZipCode		int32  `bson:"zipCode" binding:"required"`
	City		string  `bson:"city" binding:"required"`
	State 	 	string  `bson:"state" binding:"required"`
	Phone       string  `bson:"phone" binding:"required"`
	Email       string  `bson:"email" binding:"required"`
	OpenTime    string  `bson:"openTime" binding:"required"`
	CloseTime   string  `bson:"closeTime" binding:"required"`
	WorkingDays []string `bson:"workingDays" binding:"required"`
	Active	    bool    `bson:"active"`
}
 
// Returns the collection of stores from the database.
func GetAllStores() ([]Store, error) {
	collection := db.GetDBCollection("stores")
	ctx := context.TODO()

	var stores []Store
	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil { return stores, err }

	err = cursor.All(ctx, &stores)
	if err != nil { return stores, err }

	return stores, nil
}

// Deletes a store from the database.
func DeleteOneStore(id string) error {
	collection := db.GetDBCollection("stores")
	ctx := context.TODO()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil { return err }

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil { return err }

	return nil
}

// Creates a store in the database.
func CreateOneStore(store *Store) error {
	collection := db.GetDBCollection("stores")
	ctx := context.TODO()

	existingStore, err := utils.StoreExistsInDb(ctx, collection, store.Name)
	if err != nil { return err }
	if existingStore { return errors.New("store already exists") }

	_, err = collection.InsertOne(ctx, store)
	if err != nil { return err }

	return nil
}