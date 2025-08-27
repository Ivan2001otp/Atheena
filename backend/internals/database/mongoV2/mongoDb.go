package database

import (
	"context"
	"log"
	"sync"
	"time"

	envConfig "atheena/internals/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientInstanceErr error
	mongoOnce sync.Once
)

func GetMongoClient() (*mongo.Client, error) {
	MONGO_DB_PORT := envConfig.LoadEnvConfig().DB_Port

	mongoOnce.Do(func(){

		uri := "mongodb://localhost:"+MONGO_DB_PORT;
		 
		ctx, cancel := context.WithTimeout(context.Background(), 12 * time.Second)
		defer cancel()

		clientOpts := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, clientOpts)
		
		if err != nil {
			clientInstanceErr = err;
			log.Println("Something went wrong during connecting to mongo !");
			return ;
		}

		// Ping to Verfiy connection
		if err := client.Ping(ctx, nil); err != nil {
			clientInstanceErr = err;
			log.Println("Something went wrong on pinging db!")
			return;
		}

		log.Println("✅ Connected to MongoDB");
		clientInstance = client;

	})

	return clientInstance, clientInstanceErr;
}