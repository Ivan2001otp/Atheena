package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type InventoryItem struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Warehouse_Id primitive.ObjectID `json:"warehouse_id" bson:"warehouse_id"`
	Name string `json:"name" bson:"name"`
	Quantity float64 `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	Reason string `json:"reason" bson:"reason"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
	Updated_At time.Time `json:"updated_at" bson:"updated_at"`
}