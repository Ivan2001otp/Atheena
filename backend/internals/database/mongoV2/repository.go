package database

import (
	_entities "atheena/internals/entities"
	_util "atheena/internals/util"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func handleDBConnection(err error) {
	if err != nil {
	log.Fatal("Something went wrong, while DB instantiation !");
	}
}



//Sites related CRUD
func UpsertNewConstructionSiteesByAdmin(constructionSite _entities.Site) error {
	mongoDb, err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SITES);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel()

	filter := bson.M{"_id": constructionSite.ID}
	update := bson.M{
		"$set": bson.M{
            "name":    constructionSite.Name,
            "address": constructionSite.Address,
            "state":   constructionSite.State,
            "country": constructionSite.Country,
            "updated_time":    constructionSite.Updated_At,
        },
	}

	opts := options.Update().SetUpsert(true)

    result, err := collection.UpdateOne(ctx, filter, update, opts)
    if err != nil {
        log.Println("Failed to upsert construction site:", err)
        return err
    }

    if result.MatchedCount > 0 {
        log.Println("Construction site updated successfully")
    } else if result.UpsertedCount > 0 {
        log.Println("New construction site inserted with ID:", result.UpsertedID)
    }

    return nil
}





// warehouse related CRUD.
func AddWarehouseByUser(warehouse _entities.WareHouse) error {

	mongoDb, err:= GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.WAREHOUSES);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel();

	
	filter:=bson.M{"user_id":warehouse.ID}
	update := bson.M{
		"$set" : bson.M{
			 "user_id":    warehouse.User_Id,
            "name":       warehouse.Name,
            "location":   warehouse.Location,
            "address":    warehouse.Address,
            "state":      warehouse.State,
            "country":    warehouse.Country,
            "created_at": warehouse.Created_At,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("Failed to Upsert warehouse:", err);
		return err;
	}

	if result.MatchedCount > 0 {
        log.Println("Warehouse updated successfully")
    } else if result.UpsertedCount > 0 {
        log.Println("New warehouse inserted with ID:", result.UpsertedID)
    }

	return nil;
}


func DeleteWarehouseById(warehouseId primitive.ObjectID) error {
	// mongoDb, err := GetMongoClient();
	// handleDBConnection(err);

	// Fetch the inventory items of this warehouse
	return nil;
}

// Supervisor level CRUD
func UpsertNewSupervisor(supervisor _entities.Supervisor) error {
	mongoDb, err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel()
	
	filter := bson.M{"_id":supervisor.ID}
	update := bson.M{
		"$set" : bson.M{
			"admin_id":supervisor.AdminID,
			"name":supervisor.Name,
			"email":supervisor.Email,
			"phone_number":supervisor.PhoneNumber,
			"role":supervisor.Role,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Failed to upsert supervisor : ", err);
		return err;
	}


	if result.MatchedCount > 0 {
		log.Println("✅Supervisor Upserted successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅Supervisor Inserted successfully")
	}
	
	return nil;
}

func DeleteSupervisor(supervisorID primitive.ObjectID, adminID primitive.ObjectID) error {
	mongoDb, err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);

	defer cancel();

	
	result, err := collection.DeleteOne(ctx, bson.M{"_id":supervisorID, "admin_id":adminID})
	if err != nil {
		log.Println("Failed to delete supervisor : ", err);
		return err;
	}

	if result.DeletedCount > 0 {
		log.Println("✅ Supervisor deleted successfully !");
	} else {
		log.Println("⚠️ No supervisor found with the given ID");
	}

	return nil;
}


// Admin level CRUD and token operations

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

// Upserts the auth token
func InsertAuthToken(authToken _entities.AuthToken) error {
	mongoDb, err := GetMongoClient();

	handleDBConnection(err);
	
	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	// upsert query
	filter := bson.M{"user_id":authToken.User_Id}
	update := bson.M{
		"$set":bson.M{
			"email":authToken.Email,
			"role":authToken.Role,
			"refresh_token":authToken.Refresh_Token,
			"expiry_time":authToken.Expiry_Time,
			"created_at":authToken.Created_At,
		},
	}

	opts := options.Update().SetUpsert(true);

	_, err = collection.UpdateOne(ctx, filter, update, opts);
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

	log.Println("The created time is ", _util.FormatDateTime(token.Created_At))
	log.Println("The expiry time is ", _util.FormatDateTime(token.Expiry_Time))
	log.Println("Successfully fetched the refresh token ✅ - ",token.Refresh_Token);
	
	return &token, nil;
}


func DeleteUserById(objectId primitive.ObjectID) error {
	mongoDb, err := GetMongoClient();
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	result, err := collection.DeleteOne(ctx, bson.M{"_id":objectId});
	if (err != nil) {
		log.Println("Could not delete user by Id");
		log.Println(err.Error());
		return err;
	}

	log.Println("The user record delete count is ", result.DeletedCount);
	if (result.DeletedCount > 0) {
		log.Println("Successfully deleted the User from USER table.")
	} else {

		log.Println("⚠️ Failed to delete the User from USER table.");

		return fmt.Errorf("The user is already been deleted or record does not exist to delete.")
	}

	return nil;
}



func DeleteLoggedOutRefreshToken(email, role string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err);

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS);
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
	defer cancel();

	result, err := collection.DeleteOne(ctx,
		 bson.M{"email":email, "role":role}, 
		);

	log.Println("Deleted logout count : ", result.DeletedCount);

	if err != nil {
		log.Println("❌ Could not delete the Refresh token from TOKENS table.");
		log.Println(err.Error());
		return err;
	}


	if (result.DeletedCount > 0) {
		log.Println("✅ Succesfully deleted the refresh-token")
	} else {
		log.Println("❌ Could not delete the record from TOKENS table")
		return fmt.Errorf("⚠️ The target token is already deleted . Try to contact db admin to confirm.");
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