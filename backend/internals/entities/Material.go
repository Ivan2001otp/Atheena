package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This is used for site.
type Material struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`//p.k
	Name string `json:"name" bson:"name"` // cements, rods, sand, gravel

	Admin_Id primitive.ObjectID `json:"admin_id" bson:"admin_id"` //f.k
	Site_Id primitive.ObjectID `json:"site_id" bson:"site_id"`

	Quantity int `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	 
	Updated_Time time.Time	`json:"updated_time" bson:"updated_time"`
}