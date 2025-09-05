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
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err.Error())
		log.Println("Something went wrong while fetching inventories for the warehouse-id : ", warehouseId)
		return nil, err
	}

	defer cursor.Close(ctx)

	var inventoryList []_entities.InventoryItem
	if err = cursor.All(ctx, &inventoryList); err != nil {
		log.Println("Something went wrong while parsing inventory mongo documents to golang object.")
		log.Println(err.Error())
		return nil, err
	}

	log.Println("inventory records ", len(inventoryList))
	log.Println("✅ Fetched Inventories successfully for warehouseId - ", warehouseId.Hex())
	return inventoryList, nil
}

// Supervisor level CRUD
func UpsertNewSupervisor(supervisor _entities.Supervisor) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("supervisor id : ", supervisor.ID.Hex())
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

func FetchSupervisorById(supervisorID primitive.ObjectID) (*_entities.Supervisor, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)
	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()


	filter := bson.M{
		"_id": supervisorID,
	}

	var supervisor _entities.Supervisor
	err = collection.FindOne(ctx, filter).Decode(&supervisor)
	if err != nil {
		log.Println("Something went wrong while fetching supervisor.")
		log.Println(err.Error())
		return nil, err
	}

	log.Println("Successfully fetched supervisor.")
	return &supervisor, nil
}

func FetchAllSupervisorByAdminId(adminId primitive.ObjectID) ([]_entities.Supervisor, error) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.SUPERVISORS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{
		"admin_id": adminId,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println(err.Error())
		log.Println("Something went wrong while fetching supervisors")
		return nil, err
	}

	defer cursor.Close(ctx)

	var supervisorList []_entities.Supervisor
	if err = cursor.All(ctx, &supervisorList); err != nil {
		log.Println("Something went wrong while parsing supervisor mongo documents to golang object.")
		log.Println(err.Error())
		return nil, err
	}


	return supervisorList, nil
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

// Fetch logs. This query includes multiple joins
func FetchLogs(adminId primitive.ObjectID) (error, []_entities.LogisticsReport) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.LOGS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "admin_id", Value: adminId}}}}
	pipeline := mongo.Pipeline{
		matchStage,
		// Stage 1 : look up to Inventory
		{{Key: "$lookup", Value: bson.M{
			"from":         _util.INVENTORY,
			"localField":   "supply_id",
			"foreignField": "_id",
			"as":           "supply_details",
		}}},

		// unwind the array.
		{{
			Key: "$unwind", Value: bson.M{"path": "$supply_details", "preserveNullAndEmptyArrays": true},
		}},

		// Stage 2: Look up 'from-warehouse' details
		{{
			Key: "$lookup", Value: bson.M{
				"from":         _util.WAREHOUSES,
				"localField":   "from_warehouse_id",
				"foreignField": "_id",
				"as":           "from_warehouse_details",
			}}},

		{{Key: "$unwind", Value: bson.M{
			"path":   "$from_warehouse_details",
			"preserveNullAndEmptyArrays": true}}},

		// Stage 3 : Look up to-warehouse details
		{{Key: "$lookup", Value: bson.M{
			"from":         _util.WAREHOUSES,
			"localField":   "to_warehouse_id",
			"foreignField": "_id",
			"as":           "to_warehouse_details",
		}}},

		{{Key: "$unwind", Value: bson.M{
			"path":                       "$to_warehouse_details",
			"preserveNullAndEmptyArrays": true,
		}}},

		// // Stage 4 : Lookup the site details
		{{Key: "$lookup", Value: bson.M{
			"from":         _util.SITES,
			"localField":   "site_id",
			"foreignField": "_id",
			"as":           "site_details",
		}}},

		{{Key: "$unwind", Value: bson.M{
			"path":                       "$site_details",
			"preserveNullAndEmptyArrays": true,
		}}},

		// // Stage 5 : Project and final output the fields.
		{{Key: "$project", Value: bson.M{
			"_id":                     "$_id",
			"from_warehouse_name":     "$from_warehouse_details.name",
			"from_warehouse_location": "$from_warehouse_details.address",
			"from_warehouse_state":    "$from_warehouse_details.state",
			"from_warehouse_country":  "$from_warehouse_details.country",

			"to_destination_name": bson.M{
				"$ifNull": bson.A{"$to_warehouse_details.name", "$site_details.name"}},

			"to_destination_location": bson.M{
				"$ifNull": bson.A{"$to_warehouse_details.address",
					"$site_details.address"}},

			"to_destination_state": bson.M{
				"$ifNull": bson.A{"$to_warehouse_details.state", "$site_details.state"}},
			
			"to_destination_country": bson.M{
				"$ifNull": bson.A{"$to_warehouse_details.country", "$site_details.country"}},

			
			"is_site" : bson.M{
				"$cond" : bson.M{
					"if" : "$site_details",
					"then":true,
					"else":false,
				}},

			"supply_name":     "$supply_details.name",
			"supply_quantity": "$supply_details.quantity",
			"supply_unit":     "$supply_details.unit",
			"updated_time":    "$updated_time",
		}}},
	}

	var response []_entities.LogisticsReport
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println("Something went wrong while fetching log details")
		return err, nil
	}

	if err := cursor.All(ctx, &response); err != nil {
		log.Println("Something went wrong, while parsing response")
		return err, nil
	}
	return nil, response
}


func FetchOrders(adminId primitive.ObjectID) (error, []_entities.OrderItem) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.LOGS);
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel();


	matchStage := bson.D{{Key : "$match", Value : bson.D{{Key : "admin_id", Value:adminId}}}}

	pipeline := mongo.Pipeline{
		matchStage,

		// stage 2 : join orders collection using lookup
		{
			{Key:"$lookup", Value: bson.M{
				"from":_util.ORDERS,
				"localField":"_id",
				"foreignField":"log_id",
				"as":"order_details"}}},

		// stage 3
		{{
			Key : "$unwind", Value : bson.M{"path" : "$order_details", "preserveNullAndEmptyArrays":true},
		}},

		// stage 4
		{{
			Key : "$project", Value : bson.M{
				"order_id":"$order_details._id",
				"material_name":"$order_details.material_name",
				"quantity":"$order_details.quantity",
				"unit":"$order_details.unit",
				"order_type":"$order_details.order_type",
				"current_status":"$order_details.current_status",
				"trackers":"$order_details.trackers",
			}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline);
	if err != nil {
		log.Println("Failed to fetch order records.");
		log.Println(err.Error());
		return err, nil;
	}

	var result []_entities.OrderItem;
	if err := cursor.All(ctx, &result); err != nil {
		log.Println("Something went wrong, while parsing approval List");
		return err, nil
	}

	return nil, result;
}

func FetchAllApprovals(adminId primitive.ObjectID) (error, [] _entities.ApprovalResponse) {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS);
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel();


	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "admin_id", Value: adminId}}}}

	pipeline := mongo.Pipeline{
		matchStage,

		// stage look up supervisor collection
		{{Key : "$lookup", Value: bson.M{
			"from":_util.SUPERVISORS,
			"localField": "provider_id",
			"foreignField" : "_id",
			"as" : "supervisor_details",
		}}},

		{{
			Key : "$unwind", Value : bson.M{"path" : "$supervisor_details", "preserveNullAndEmptyArrays":true},
		}},


		// stage look up inventory collection.
		{{
			Key : "$lookup", Value : bson.M{
				"from" : _util.INVENTORY,
				"localField": "supply_id",
				"foreignField" : "_id",
				"as" : "inventory_details",
			}}},

		{{
			Key : "$unwind", Value : bson.M{"path" : "$inventory_details", "preserveNullAndEmptyArrays":true},
		}},


		// stage look up source/from-warehouse collection.
		{{
			Key : "$lookup", Value : bson.M{
				"from":_util.WAREHOUSES,
				"localField":"from_id",
				"foreignField":"_id",
				"as" : "from_warehouse_details",
			}}},
		{{
			Key : "$unwind", Value : bson.M{"path" : "$from_warehouse_details", "preserveNullAndEmptyArrays" : true}}},



		// projection stage
		{{
			Key : "$project", Value : bson.M{
				"_id": "$_id",
				"from_warehouse_name": "$from_warehouse_details.name",
				"from_warehouse_location" : "$from_warehouse_details.address",
				"from_warehouse_state":"$from_warehouse_details.state",
				"from_warehouse_country": "$from_warehouse_details.country",
				"status":"$status",
				"reason":"$reason",
				"supply_name":     "$inventory_details.name",
				"supervisor_name": "$supervisor_details.name",
				"updated_time":    "$updated_at",
		}}},
	}


	var approvalList []_entities.ApprovalResponse;
 	cursor, err :=  collection.Aggregate(ctx, pipeline);

	if err != nil {
		log.Println("Something went wrong while fetching approvals.");
		log.Println(err.Error());
		return err, nil;
	}

	if err := cursor.All(ctx, &approvalList); err != nil {
		log.Println("Something went wrong, while parsing approval List");
		return err, nil
	}

	return nil,approvalList;

}
