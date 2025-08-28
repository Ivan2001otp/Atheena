package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderTracker struct {
	Order_Status string    `json:"order_status" bson:"order_status"`
	Created_Time time.Time `json:"created_time" bson:"created_time"`
}

type Order struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Log_Id primitive.ObjectID `json:"log_id" bson:"log_id"`
	Material_Name string `json:"material_name" bson:"material_name"`
	Quantity int `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	Type string `json:"order_type" bson:"order_type"`
	Updated_Time time.Time `json:"updated_time" bson:"updated_time"`
	Order_Trackers []OrderTracker `json:"trackers" bson:"trackers"`
}