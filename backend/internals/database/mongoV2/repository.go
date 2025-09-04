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
		log.Fatal("❌Something went wrong, while DB instantiation !")
	}
}

// Approval related CRUD
func UpsertNewApproval(approvalNotification _entities.ApprovalTypeNotification) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": approvalNotification.ApprovalID,
	}

	update := bson.M{
		"$set": bson.M{
			"provider_id":    approvalNotification.ProviderID,
			"admin_id":       approvalNotification.AdminID,
			"supply_id":      approvalNotification.SupplyID,
			"from_id":        approvalNotification.FromID,
			"destination_id": approvalNotification.DestinationID,
			"reason":         approvalNotification.Reason,
			"status":         approvalNotification.Status,
			"created_at":     approvalNotification.CreatedAt,
			"updated_at":     approvalNotification.UpdatedAt,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println(err.Error())
		log.Println("❌Something went wrong while upserting approval !")
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("✅✅Approval record updated successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅New Approval record inserted with ID:", result.UpsertedID)
	}

	return nil
}

// Sites related CRUD
func UpsertNewConstructionSiteesByAdmin(constructionSite _entities.Site) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SITES)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": constructionSite.ID}
	update := bson.M{
		"$set": bson.M{
			"name":         constructionSite.Name,
			"address":      constructionSite.Address,
			"state":        constructionSite.State,
			"country":      constructionSite.Country,
			"user_id":      constructionSite.AdminId,
			"updated_time": constructionSite.Updated_At,
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Failed to upsert construction site:", err)
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("✅✅Construction site updated successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅New construction site inserted with ID:", result.UpsertedID)
	}

	return nil
}

func FetchSitesbyAdminId(adminId primitive.ObjectID) ([]_entities.Site, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SITES)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": adminId}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("❌something went wrong while fetching all wareshouses for the given adminId")
		return nil, err
	}

	defer cursor.Close(ctx)

	var siteList []_entities.Site
	if err = cursor.All(ctx, &siteList); err != nil {
		log.Println("❌something went wrong while parsing construction sites list.")
		return nil, err
	}

	return siteList, nil
}

// warehouse related CRUD.
func AddWarehouseByUser(warehouse _entities.WareHouse) error {

	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.WAREHOUSES)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": warehouse.ID}
	update := bson.M{
		"$set": bson.M{
			"user_id":    warehouse.User_Id,
			"name":       warehouse.Name,
			"pin":        warehouse.Pin,
			"address":    warehouse.Address,
			"state":      warehouse.State,
			"country":    warehouse.Country,
			"created_at": warehouse.Created_At,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Failed to Upsert warehouse:", err)
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("✅✅Warehouse updated successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅New warehouse inserted with ID:", result.UpsertedID)
	}

	return nil
}

func FetchWarehouseById(adminId primitive.ObjectID) ([]_entities.WareHouse, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.WAREHOUSES)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": adminId}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("❌something went wrong while fetching all wareshouses for the given adminId")
		return nil, err
	}

	defer cursor.Close(ctx)

	var warehouseList []_entities.WareHouse
	if err = cursor.All(ctx, &warehouseList); err != nil {
		log.Println("❌something went wrong while parsing warehouse list.")
		return nil, err
	}

	return warehouseList, nil
}

func DeleteWarehouseById(warehouseId primitive.ObjectID) error {
	// mongoDb, err := GetMongoClient();
	// handleDBConnection(err);

	// Fetch the inventory items of this warehouse
	return nil
}

// Inventory related CRUD.
func AddNewInventoryItems(inventoryItem _entities.InventoryItem) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": inventoryItem.ID}
	update := bson.M{
		"$set": bson.M{
			"warehouse_id": inventoryItem.Warehouse_Id,
			"name":         inventoryItem.Name,
			"quantity":     inventoryItem.Quantity,
			"unit":         inventoryItem.Unit,
			"reason":       inventoryItem.Reason,
			"created_at":   inventoryItem.Created_At,
			"updated_at":   inventoryItem.Updated_At,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Could not upsert the inventory Item.")
	}

	if result.MatchedCount > 0 {
		log.Println("✅✅Supervisor Upserted successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅Supervisor Inserted successfully")
	}

	return nil
}

func FetchInventoryByWarehouseId(warehouseId primitive.ObjectID) ([]_entities.InventoryItem, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"warehouse_id": warehouseId}
	cursor,err := collection.Find(ctx, filter);
	if err != nil {
		log.Println(err.Error());
		log.Println("Something went wrong while fetching inventories for the warehouse-id : ",warehouseId);
		return nil,err;
	}

	defer cursor.Close(ctx);

	var inventoryList []_entities.InventoryItem;
	if err = cursor.All(ctx, &inventoryList); err != nil {
		log.Println("Something went wrong while parsing inventory mongo documents to golang object.");
		log.Println(err.Error())
		return nil, err;
	}

	log.Println("inventory records ", len(inventoryList));
	log.Println("✅ Fetched Inventories successfully for warehouseId - ", warehouseId.Hex());
	return inventoryList, nil;
}

// Supervisor level CRUD
func UpsertNewSupervisor(supervisor _entities.Supervisor) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": supervisor.ID}
	update := bson.M{
		"$set": bson.M{
			"admin_id":     supervisor.AdminID,
			"name":         supervisor.Name,
			"email":        supervisor.Email,
			"phone_number": supervisor.PhoneNumber,
			"role":         supervisor.Role,
		},
	}

	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Failed to upsert supervisor : ", err)
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("✅✅Supervisor Upserted successfully")
	} else if result.UpsertedCount > 0 {
		log.Println("✅Supervisor Inserted successfully")
	}

	return nil
}

func DeleteSupervisor(supervisorID primitive.ObjectID, adminID primitive.ObjectID) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": supervisorID, "admin_id": adminID})
	if err != nil {
		log.Println("❌Failed to delete supervisor : ", err)
		return err
	}

	if result.DeletedCount > 0 {
		log.Println("✅ Supervisor deleted successfully !")
	} else {
		log.Println("⚠️ No supervisor found with the given ID")
	}

	return nil
}

func FetchSupervisorById(supervisorID primitive.ObjectID) (*_entities.Supervisor,error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{
		"_id":supervisorID,
	}

	var supervisor _entities.Supervisor
	err = collection.FindOne(ctx, filter).Decode(&supervisor);
	if err != nil {
		log.Println("Something went wrong while fetching supervisor.");
		log.Println(err.Error());
		return nil,err;
	}

	log.Println("Successfully fetched supervisor.");
	return &supervisor, nil;
}

func FetchAllSupervisorByAdminId(adminId primitive.ObjectID) ([]_entities.Supervisor ,error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{
		"admin_id":adminId,
	}

	cursor,err := collection.Find(ctx, filter);
	if err != nil {
		log.Println(err.Error());
		log.Println("Something went wrong while fetching supervisors");
		return nil,err;
	}

	defer cursor.Close(ctx);

	var supervisorList []_entities.Supervisor;
	if err = cursor.All(ctx, &supervisorList); err != nil {
		log.Println("Something went wrong while parsing supervisor mongo documents to golang object.");
		log.Println(err.Error())
		return  nil,err;
	}


	return supervisorList, nil;
} 

// Admin level CRUD and token operations
func InsertNewUser(user _entities.User) error {
	mongoDb, err := GetMongoClient()

	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// insert document...
	_, err = collection.InsertOne(ctx, user)
	return err
}

// Upserts the auth token
func InsertAuthToken(authToken _entities.AuthToken) error {
	mongoDb, err := GetMongoClient()

	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// upsert query
	filter := bson.M{"user_id": authToken.User_Id}
	update := bson.M{
		"$set": bson.M{
			"email":         authToken.Email,
			"role":          authToken.Role,
			"refresh_token": authToken.Refresh_Token,
			"expiry_time":   authToken.Expiry_Time,
			"created_at":    authToken.Created_At,
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func GetTokenByRefreshToken(refreshToken string) (*_entities.AuthToken, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	var token _entities.AuthToken
	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&token)
	if err != nil {
		log.Println(err.Error())
		return &token, err
	}

	log.Println("The created time is ", _util.FormatDateTime(token.Created_At))
	log.Println("The expiry time is ", _util.FormatDateTime(token.Expiry_Time))
	log.Println("Successfully fetched the refresh token ✅ - ", token.Refresh_Token)

	return &token, nil
}

func DeleteUserById(objectId primitive.ObjectID) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		log.Println("❌Could not delete user by Id")
		log.Println(err.Error())
		return err
	}

	log.Println("The user record delete count is ", result.DeletedCount)
	if result.DeletedCount > 0 {
		log.Println("✅Successfully deleted the User from USER table.")
	} else {

		log.Println("⚠️ Failed to delete the User from USER table.")

		return fmt.Errorf("The user is already been deleted or record does not exist to delete.")
	}

	return nil
}

func DeleteLoggedOutRefreshToken(email, role string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.TOKENS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("Email is %s, role is %s \n", email, role)
	result, err := collection.DeleteOne(ctx,
		bson.M{"email": email, "role": role},
	)

	log.Println("Deleted logout count : ", result.DeletedCount)
	if err != nil {
		log.Println("❌ Could not delete the Refresh token from TOKENS table.")
		log.Println(err.Error())
		return err
	}

	if result.DeletedCount > 0 {
		log.Println("✅ Succesfully deleted the refresh-token")
	} else {
		log.Println("❌ Could not delete the record from TOKENS table")
		return fmt.Errorf("⚠️ The target token is already deleted . Try to contact db admin to confirm.")
	}

	return nil
}

func EmailExists(email string) (*_entities.User, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.USERS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user _entities.User
	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// email does not exists.
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
