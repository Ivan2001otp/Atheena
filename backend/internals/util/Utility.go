package util

import (
	"log"
	"sync"

	"github.com/google/uuid"
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