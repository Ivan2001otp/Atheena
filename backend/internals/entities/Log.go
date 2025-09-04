package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogType string

const (
	InboundLog LogType = "INBOUND"
	OutboundLog LogType = "OUTBOUND"
	SiteLog LogType = "CONSTRUCTION_SITE"
)
type CustomLog struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"` // p.k
	AdminId primitive.ObjectID `json:"admin_id" bson:"admin_id"`
	SupplyID primitive.ObjectID `json:"supply_id" bson:"supply_id"`// f.k

	Supervisor_ID primitive.ObjectID `json:supervisor_id" bson:"supervisor_id"`
	From_Warehouse_Id primitive.ObjectID `json:"from_warehouse_id"  bson:"from_warehouse_id"`
	To_Warehouse_Id primitive.ObjectID `json:"to_warehouse_id" bson:"to_warehouse_id"`
	
	Site_Id primitive.ObjectID `json:"site_id" bson:"site_id"`
	Log_Type LogType `json:"log_type" bson:"log_type"`
	Updated time.Time `json:"updated_time" bson:"updated_time"`
}

type LogisticsReport struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	FromWarehouseName      string             `bson:"from_warehouse_name"`
	FromWarehouseLocation  string             `bson:"from_warehouse_location"`
	ToDestinationName      string             `bson:"to_destination_name"`
	ToDestinationLocation  string             `bson:"to_destination_location"`
	SupplyName             string             `bson:"supply_name"`
	SupplyQuantity         float64            `bson:"supply_quantity"`
	SupplyUnit             string             `bson:"supply_unit"`
	UpdatedTime            time.Time          `bson:"updated_time"`
}