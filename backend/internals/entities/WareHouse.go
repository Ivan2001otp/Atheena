package entities

import ("go.mongodb.org/mongo-driver/bson/primitive"
"time"
)

type WareHouse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	User_Id    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name       string             `json:"name" bson:"name"`
	Location   string             `json:"location" bson:"location"`
	Address    string             `json:"address" bson:"address"`
	State      string             `json:"state" bson:"state"`
	Country    string             `json:"country" bson:"country"`
	Created_At time.Time          `json:"created_at" bson:"created_at"`
}
