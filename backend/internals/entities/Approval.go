package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApprovalResponse struct {
	ID                      primitive.ObjectID `bson:"_id" json:"_id"`
	From_Warehouse_Name     string             `bson:"from_warehouse_name" json:"from_warehouse_name"`
	From_Warehouse_Location string             `bson:"from_warehouse_location" json:"from_warehouse_location"`

	From_Warehouse_State   string `bson:"from_warehouse_state" json:"from_warehouse_state"`
	From_Warehouse_Country string `bson:"from_warehouse_country" json:"from_warehouse_country"`
	Status                 string `bson:"status" json:"status"`
	Reason                 string `bson:"reason" json:"reason"`

	Supply_Name     string    `bson:"supply_name" json:"supply_name"`
	Supervisor_Name string    `bson:"supervisor_name" json:"supervisor_name"`
	Updated_Time    time.Time `bson:"updated_time" json:"updated_time"`
}
