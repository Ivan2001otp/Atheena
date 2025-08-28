package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Material struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`//p.k
	Name string `json:"name" bson:"name"` // cements, rods, sand, gravel

	WareHouse_Id primitive.ObjectID `json:"warehouse_id" bson:"warehouse_id"`
	User_Id primitive.ObjectID `json:"user_id" bson:"user_id"` //f.k

	Reason string `json:"reason" bson:"reason"`
	Quantity int `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	 
	Created_Time time.Time `json:"created_time" bson:"created_time"`
	Updated_Time time.Time	`json:"updated_time" bson:"updated_time"`
}