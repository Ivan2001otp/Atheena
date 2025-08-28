package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Quantity int `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	Created_Time time.Time `json:"created_time" bson:"created_time"`
}