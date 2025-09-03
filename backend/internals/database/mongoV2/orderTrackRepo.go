package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

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

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.LOGS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": logRecord.ID,
	}

	updateTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	update := bson.M{
		"$set": bson.M{
			"admin_id":          logRecord.AdminId,
			"supply_id":         logRecord.SupplyID,
			"supervisor_id":     logRecord.Supervisor_ID,
			"from_warehouse_id": logRecord.From_Warehouse_Id,
			"to_warehouse_id":   logRecord.To_Warehouse_Id,
			"site_id":           logRecord.Site_Id,
			"updated_time":      updateTime,
		},
	}

	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println("❌Something went wrong while upserting log record.")
		log.Println(err.Error())
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("The target log record found.")
	}

	if result.ModifiedCount > 0 {
		log.Println("The target log record was successfully updated.")
	}

	return nil
}

func UpdateApprovalStatus(approvalId primitive.ObjectID, status string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": approvalId,
	}

	updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": updated_time,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println(err.Error())
		log.Println("❌Failed to update the approval record .")
		return err
	}

	if result.MatchedCount > 0 {
		log.Println("The target approval record found.")
	}

	if result.ModifiedCount > 0 {
		log.Println("The target approval record was successfully updated.")
	}

	return nil
}

// update the quantity of inventory of the given warehouse
func UpdateInventoryQuantity(warehouseId primitive.ObjectID, inventoryId primitive.ObjectID, quantity float64) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var inv struct {
		Quantity float64 `bson:"quantity"`
	}

	err = collection.FindOne(ctx, bson.M{"_id": inventoryId, "warehouse_id": warehouseId}).Decode(&inv)
	if err != nil {
		log.Println(err.Error())
		log.Println("❌Something went wrong while fetching inventory for the given _id and warehouse_id or record not found.")
		return err
	}

	if inv.Quantity-quantity < 0 {
		return fmt.Errorf("insufficient stock: current=%.2f, requested=%.2f", inv.Quantity, quantity)
	}

	filter := bson.M{
		"_id":          inventoryId,
		"warehouse_id": warehouseId,
	}

	updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	update := bson.M{
		"$inc": bson.M{
			"quantity": -quantity,
		},
		"$set": bson.M{
			"updated_at": updated_time,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Println(err.Error())
		log.Println("❌Failed to update the approval record .")
		return err
	}

	if result.ModifiedCount > 0 {
		log.Println("The target approval record was successfully updated.")
	} else {
		log.Println("No Inventory document updated.")
	}

	return nil
}

func CreateOrderRecord(logId primitive.ObjectID, quantity float64, orderType string, materialName string, unit string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.ORDERS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order _entities.Order

	order.ID = primitive.NewObjectID()
	order.Log_ID = logId
	order.Quantity = quantity
	order.Type = _entities.OrderType(orderType)
	order.Material_Name = materialName
	order.Current_Status = _entities.OrderPlaced

	// confirmed, shipped, delivered, completed.
	var orderTracker []_entities.OrderTracker
	created_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	orderTracker = append(orderTracker, _entities.OrderTracker{Order_Status: _entities.OrderPlaced, Created_Time: created_time})

	order.Order_Trackers = orderTracker
	order.Unit = unit
	order.Updated_Time, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Println("❌Something went wrong while inserting order item.")
		log.Println(err.Error())
		return err
	}

	log.Println("Order is created successfully - ", result.InsertedID)
	return nil
}

func UpdateOrder(orderId primitive.ObjectID, status string) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	collection := mongoDb.Database(_util.DATABASE).Collection(_util.ORDERS)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": orderId,
	}

	updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	tracker := bson.M{
		"order_status": status,
		"created_time": updated_time,
	}

	update := bson.M{
		"$set": bson.M{"current_status": status, "updated_time": updated_time},
		"$push": bson.M{
			"trackers": tracker,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("failed to update order: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no order found with id %v", orderId.Hex())
	}

	log.Printf("Order %v updated successfully with status: %s\n", orderId.Hex(), status)
	return nil
}

// Uses mongo transaction to fire multi query to maintain atomicity.
func UpdateFinalOrderStatus(status string, materialName string, orderId primitive.ObjectID, toWarehouseID primitive.ObjectID, siteID primitive.ObjectID, approvalID primitive.ObjectID) error {
	mongoDb, err := GetMongoClient()
	handleDBConnection(err)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	
	session, err := mongoDb.StartSession()
	if err != nil {
		log.Println("Failed to start session for mongo transation (UpdateFinalOrderStatus)")
		return err
	}

	defer session.EndSession(ctx)

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		orderCollection := mongoDb.Database(_util.DATABASE).Collection(_util.ORDERS)
		inventoryCollection := mongoDb.Database(_util.DATABASE).Collection(_util.INVENTORY)
		logCollection := mongoDb.Database(_util.DATABASE).Collection(_util.LOGS)
		approvalCollection := mongoDb.Database(_util.DATABASE).Collection(_util.APPROVALS)
		materialCollection := mongoDb.Database(_util.DATABASE).Collection(_util.MATERIALS)

		// update the order
		filter := bson.M{
			"_id": orderId,
		}

		updated_time, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		tracker := bson.M{
			"order_status": _entities.OrderStatus(status),
			"created_time": updated_time,
		}

		update := bson.M{
			"$set": bson.M{"current_status": _entities.OrderStatus(status), "updated_time": updated_time},
			"$push": bson.M{
				"trackers": tracker,
			},
		}

		_, err := orderCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Println(err.Error())
			log.Println("Failed to update the order status")
			return nil, err
		}
		log.Println("Updated the order successfully ✅ ");

		// get the order quantity
		var order _entities.Order
		err = orderCollection.FindOne(ctx, bson.M{"_id": orderId}).Decode(&order)
		if err != nil {
			log.Println("Failed to fetch order to continue udpates.")
			log.Println(err.Error())
			return nil, err
		}
		log.Println("Fetched order successfully ✅ ");

		var customLog _entities.CustomLog
		err = logCollection.FindOne(ctx, bson.M{"_id": order.Log_ID}).Decode(&customLog)
		if err != nil {
			log.Println(err.Error())
			log.Println("Failed to fetch logs.")
			return nil, err
		}
		log.Println("Fetched the logs ✅ ");


		// update the inventory of target site/warehouse.
		if siteID == primitive.NilObjectID {
			// update the inventory of warehouse.
			updated_time, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			filter = bson.M{
				"name":         order.Material_Name,
				"warehouse_id": toWarehouseID,
			}

			update = bson.M{
				"$inc": bson.M{
					"quantity": order.Quantity,
				},
				"$set": bson.M{
					"updated_at": updated_time,
				},
			}

			_, err = inventoryCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				log.Println(err.Error())
				log.Println("Failed to update inventory of target warehouse.")
				return nil, err
			}
			log.Println("✅ Updated the inventory at warehouse.");
		} else {
			// update the quantity of site.

			// fetch the admin id from log_id
			updated_time, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			// update on material collection.
			filter = bson.M{
				"name":     materialName,
				"admin_id": customLog.AdminId,
				"site_id":  siteID,
			}

			update := bson.M{
				"$inc": bson.M{
					"quantity": order.Quantity, // increment if exists
				},
				"$set": bson.M{
					"updated_time": updated_time,
					"unit":         order.Unit, // keep unit updated
				},
				"$setOnInsert": bson.M{
					"_id":          primitive.NewObjectID(),
					"name":         materialName,
					"admin_id":     customLog.AdminId,
					"site_id":      siteID,
				},
			}

			opts := options.Update().SetUpsert(true)

			_, err = materialCollection.UpdateOne(ctx, filter, update, opts)
			if err != nil {
				log.Println(err.Error())
				return nil, fmt.Errorf("failed to upsert material: %v", err)
			}
			log.Println("✅ Updated the inventory at construction site.");
		}

		// update approval status to completed.

		filter = bson.M{
			"_id":      approvalID,
			"admin_id": customLog.AdminId,
		}

		var approval_status string = ""
		if status == "DELIVERED" {
			approval_status = "completed"
		}
		update = bson.M{
			"$set": bson.M{
				"status": approval_status,
			},
		}

		res, err := approvalCollection.UpdateOne(ctx, filter, update)

		if err != nil {
			log.Println("Could not update the approval status")
			log.Println(err.Error())
			return nil, err
		}

		if res.ModifiedCount > 0 {
			log.Println("Updated the approval successfully")			
		}

		return nil, nil
	}

	

	_, err = session.WithTransaction(ctx, callback)
	return err

}
