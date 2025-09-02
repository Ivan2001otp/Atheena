package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_entities "atheena/internals/entities"
	_util "atheena/internals/util"
	"context"
	"log"
	"time"
)


func UpsertLog(logRecord _entities.CustomLog) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": logRecord.ID,
	}

	updateTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

 	update := bson.M{
		"$set" : bson.M {
		"admin_id":logRecord.AdminId,
		"supply_id":logRecord.SupplyID,
		"supervisor_id":logRecord.Supervisor_ID,
		"from_warehouse_id": logRecord.From_Warehouse_Id,
		"to_warehouse_id":logRecord.To_Warehouse_Id,
		"site_id": logRecord.Site_Id,
		"updated_time": updateTime,
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, opts);
	if err != nil {
		log.Println("❌Something went wrong while upserting log record.");
		log.Println(err.Error());
		return err;
	}


	if (result.MatchedCount > 0){
		log.Println("The target log record found.");
	}

	if (result.ModifiedCount > 0) {
		log.Println("The target log record was successfully updated.");
	}

	return nil;
}

func UpdateApprovalStatus(approvalId primitive.ObjectID, status string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id" : approvalId,
	}

	updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{
		"$set" : bson.M{
			"status" : status,
			"updated_at": updated_time,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if (err != nil ){
		log.Println(err.Error());
		log.Println("❌Failed to update the approval record .");
		return err;
	}

	if (result.MatchedCount > 0){
		log.Println("The target approval record found.");
	}

	if (result.ModifiedCount > 0) {
		log.Println("The target approval record was successfully updated.");
	}

	return nil;
}


// update the quantity of inventory of the given warehouse
func UpdateInventoryQuantity(warehouseId primitive.ObjectID, inventoryId primitive.ObjectID, quantity float64) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY);
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel();

	var inv struct {
		Quantity float64 `bson:"quantity"`
	}

	err = collection.FindOne(ctx, bson.M{"_id":inventoryId, "warehouse_id":warehouseId}).Decode(&inv);
	if err != nil {
		log.Println(err.Error());
		log.Println("❌Something went wrong while fetching inventory for the given _id and warehouse_id or record not found.");
		return err;
	}

	if (inv.Quantity - quantity < 0) {
		return fmt.Errorf("insufficient stock: current=%.2f, requested=%.2f", inv.Quantity, quantity);
	}

	filter := bson.M{
		"_id" : inventoryId,
		"warehouse_id": warehouseId,
	}

	updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{
		"$inc" : bson.M{
		  "quantity":-quantity,
		},
		"$set" : bson.M{
			"updated_at": updated_time,
		},
	}
	
	result, err := collection.UpdateOne(ctx, filter, update)

	if (err != nil ){
		log.Println(err.Error());
		log.Println("❌Failed to update the approval record .");
		return err;
	}


	if (result.ModifiedCount > 0) {
		log.Println("The target approval record was successfully updated.");
	} else {
		log.Println("No Inventory document updated.")
	}

	return nil;
}


func CreateOrderRecord(logId primitive.ObjectID, quantity float64, orderType string, materialName string, unit string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY);
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second);
	defer cancel();


	var order _entities.Order;

	order.ID = primitive.NewObjectID()
	order.Log_ID = logId;
	order.Quantity = quantity;
	order.Type = _entities.OrderType(orderType)
	order.Material_Name = materialName;

	// confirmed, shipped, delivered, completed.
	var orderTracker []_entities.OrderTracker;
	created_time,_ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
	orderTracker = append(orderTracker, _entities.OrderTracker{Order_Status: "confirmed", Created_Time: created_time});

	order.Order_Trackers = orderTracker
	order.Unit = unit;
	order.Updated_Time,_  = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

	result, err := collection.InsertOne(ctx, order);
	if err != nil {
		log.Println("❌Something went wrong while inserting order item.");
		log.Println(err.Error());
		return err;
	}

	log.Println("Order is created successfully - ",result.InsertedID);
	return nil;
}