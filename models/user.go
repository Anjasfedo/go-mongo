package models

import (
	"gopkg.in/mgo.v2/bson" // Importing bson package for working with BSON data
)

// User represents the model for user data
type User struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`       // User's unique identifier
	Name   string        `json:"name" bson:"name"`     // User's name
	Gender string        `json:"gender" bson:"gender"` // User's gender
	Age    int           `json:"age" bson:"age"`       // User's age
}
