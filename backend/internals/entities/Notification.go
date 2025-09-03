package entities

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)


type NotificationStatus string

const (
	StatusPending NotificationStatus = "pending"
	StatusApproved NotificationStatus = "approved"
	StatusDismissed NotificationStatus = "dismissed"
)

type ApprovalTypeNotification struct {
	ApprovalID     primitive.ObjectID `json:"_id" bson:"_id"`
    ProviderID     primitive.ObjectID `json:"provider_id" bson:"provider_id"`
    AdminID        primitive.ObjectID `json:"admin_id" bson:"admin_id"`
    SupplyID       primitive.ObjectID `json:"supply_id" bson:"supply_id"`
    FromID         primitive.ObjectID `json:"from_id" bson:"from_id"`
    DestinationID  primitive.ObjectID `json:"destination_id" bson:"destination_id"`
    Reason         string            `json:"reason" bson:"reason"`
    Status         NotificationStatus `json:"status" bson:"status"`
    CreatedAt      time.Time         `json:"created_at" bson:"created_at"`
    UpdatedAt      time.Time         `json:"updated_at" bson:"updated_at"`
}
