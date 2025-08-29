package util

import (
	"log"
	"sync"
	"time"
	"strings"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	uuid_generator uuid.UUID
	mu sync.Mutex
)

func GenerateRandomUUID() string {
	mu.Lock()
	defer mu.Unlock();
	
	uuid_generator = uuid.New()
	log.Println("Generated new UUID : ", uuid_generator.String())
	return uuid_generator.String();
}

func GenerateObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func GenerateCreateDateTime()( time.Time, error) {
	return time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
}


func ToUpper(str string) string {
	return strings.ToUpper(str);
}

func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 03:04:05 PM");
}