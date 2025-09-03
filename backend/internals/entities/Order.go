package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type OrderType string;
type OrderStatus string;

const (
	OrderPlaced OrderStatus = "ORDER_PLACED"
	OutForDelivery OrderStatus = "OUT_FOR_DELIVERY"
	Delivered OrderStatus = "DELIVERED"
)

const (
	InboundOrder OrderType = "INBOUND"
	OutboundOrder OrderType = "OUTBOUND"
	ConstructionSiteOrder OrderType = "CONSTRUCTION_SITE"
)

type OrderTracker struct {
	Order_Status OrderStatus    `json:"order_status" bson:"order_status"`
	Created_Time time.Time `json:"created_time" bson:"created_time"`
}

type Order struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Log_ID primitive.ObjectID `json:"log_id" bson:"log_id"`
	Material_Name string `json:"material_name" bson:"material_name"`
	Quantity float64 `json:"quantity" bson:"quantity"`
	Unit string `json:"unit" bson:"unit"`
	Type OrderType `json:"order_type" bson:"order_type"`
	Current_Status OrderStatus `json:"current_status" bson:"current_type"`
	Updated_Time time.Time `json:"updated_time" bson:"updated_time"`
	Order_Trackers []OrderTracker `json:"trackers" bson:"trackers"`
}