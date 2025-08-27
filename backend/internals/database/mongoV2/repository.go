package database

import (
	_entities "atheena/internals/entities"
	_util "atheena/internals/util"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func handleDBConnection(err error) {
	if err != nil {
	log.Fatal("Something went wrong, while DB instantiation !");
	}
}

func InsertNewUser(user _entities.User) error {
	mongoDb, err := GetMongoClient()

	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);

	defer cancel();

	// insert document...
	_, err = collection.InsertOne(ctx, user);
	return err;
} 

func InsertAuthToken(authToken _entities.AuthToken) error {
	mongoDb, err := GetMongoClient();

	handleDBConnection(err);
	
	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);

	defer cancel();

	// insert auth-token document.
	_, err = collection.InsertOne(ctx, authToken);
	return err;
}

func GetTokenByRefreshToken(refreshToken string) (*_entities.AuthToken, error ) {
	mongoDb , err := GetMongoClient()
	handleDBConnection(err);

	var token _entities.AuthToken;

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);

	ctx,cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel();

	if err != nil {
		log.Println(err.Error())
		return nil, err;
	}

	err = collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&token);
	if err != nil {
		log.Println(err.Error());
		return &token, err;
	}

	return &token, nil;
}