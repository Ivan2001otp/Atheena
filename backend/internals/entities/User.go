package entities

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



type Supervisor struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	AdminID primitive.ObjectID `json:"admin_id" bson:"admin_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Role string `json:"role" bson:"role"`
}

type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Role string `json:"role" bson:"role"`// ADMIN or SUPERVISOR
	Visited_Time time.Time `json:"visited_time" bson:"visited_time"`
}