package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive")

// This the model representation for resource consumption
type Resource struct {
	ID primitive.ObjectID `json:"id" bson:"id"`//p.k
	WareHouse_Id primitive.ObjectID `json:"warehouse_id" bson:"warehouse_id"`
	Site_Id primitive.ObjectID `json:"site_id" bson:"site_id"`
	Name string `json:"name" bson:"name"` // resource name like rods, gravel... etc
	Unit string `json:"unit" bson:"unit"`
	Quantity int `json:"quantity" bson:"quantity"`
	Reason string `json:"reason" bson:"reason"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}