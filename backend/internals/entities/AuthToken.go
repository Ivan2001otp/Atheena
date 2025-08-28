package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive")

type AuthToken struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	User_Id primitive.ObjectID `json:"user_id" bson:"user_id"`
	Email string `json:"email" bson:"email"`
	Role string `json:"role" bson:"role"`
	Refresh_Token string `json:"refresh_token" bson:"refresh_token"`
	Expiry_Time time.Time `json:"expiry_time" bson:"expiry_time"`
	Created_At time.Time `json:"created_time" bson:"created_at"`
}