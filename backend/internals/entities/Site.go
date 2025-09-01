package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Site struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	AdminId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	State string `json:"state" bson:"state"`
	Country string `json:"country" bson:"country"`
	Updated_At time.Time `json:"updated_time" bson:"updated_time"`
}