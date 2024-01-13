package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
