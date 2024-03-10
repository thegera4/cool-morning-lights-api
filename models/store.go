package models

import (

)

// Product struct to represent a store in the application.
type Store struct {
	Id          string  `bson:"_id,omitempty"`
	Name        string  `bson:"name" binding:"required"`
	Address    	string  `bson:"location" binding:"required"`
	ZipCode		string  `bson:"zipCode" binding:"required"`
	City		string  `bson:"city" binding:"required"`
	State 	 	string  `bson:"state" binding:"required"`
	Phone       string  `bson:"phone" binding:"required"`
	Email       string  `bson:"email" binding:"required"`
	OpenTime    string  `bson:"openTime" binding:"required"`
	CloseTime   string  `bson:"closeTime" binding:"required"`
	WorkingDays []string `bson:"workingDays" binding:"required"`
	Active	    bool    `bson:"active"`
}