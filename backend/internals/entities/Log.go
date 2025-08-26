package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CustomLog struct {
	ID primitive.ObjectID `json:"id" bson:"id"` // p.k
	Supply_Id primitive.ObjectID `json:"supply_id" bson:"supply_id"`// f.k
	From_Warehouse_Id primitive.ObjectID `json:"from_warehouse_id"  bson:"from_warehouse_id"`
	To_Warehouse_Id primitive.ObjectID `json:"to_warehouse_id" bson:"to_warehouse_id"`
	Site_Id primitive.ObjectID `json:"site_id" bson:"site_id"`
	Log_Type string `json:"log_type" bson:"log_type"`
	Updated time.Time `json:"updated_time" bson:"updated_time"`
}