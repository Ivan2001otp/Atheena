package database

import (
	_entities "atheena/internals/entities"
	_util "atheena/internals/util"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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



func DeleteUserById(objectId primitive.ObjectID) error {
	mongoDb, err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	result, err := collection.DeleteOne(ctx, bson.M{"id":objectId});
	if (err != nil) {
		log.Println("Could not delete user by Id");
		log.Println(err.Error());
		return err;
	}

	
	if (result.DeletedCount > 0) {
		log.Println("Successfully deleted the User from USER table.")
	} else {
		log.Println("Failed to delete the User from USER table.");
	}

	return nil;
}



func DeleteLoggedOutRefreshToken(email, role string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	result, err := collection.DeleteOne(ctx, bson.M{"email":email, "role":role});

	if err != nil {
		log.Println("❌ Could not delete the Refresh token from TOKENS table.");
		log.Println(err.Error());
		return err;
	}


	if (result.DeletedCount > 0) {
		log.Println("✅ Succesfully deleted the refresh-token")
	} else {
		log.Println("❌ Could not delete the record from TOKENS table")
	}

	return nil;
}


func EmailExists(email string) (*_entities.User, error) {

	mongoDb , err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	var user _entities.User
	err = collection.FindOne(ctx, bson.M{"email":email}).Decode(&user);

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// email does not exists.
			return nil, nil;
		}

		return nil, err;
	}

	return &user, nil;
}