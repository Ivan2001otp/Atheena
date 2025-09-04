package handlers

import (
	_mongo "atheena/internals/database/mongoV2"
	_websockets "atheena/internals/database/websockets"
	_entities "atheena/internals/entities"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AskForApproval(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("supposed to be POST request")
		http.Error(w, "Supposed to be POST request.", http.StatusBadRequest)
		return
	}

	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println(err.Error())
		log.Println("❌Failed to parse the request body.")
		return
	}

	var payload _entities.ApprovalTypeNotification
	payload.ApprovalID = primitive.NewObjectID()

	supervisorIdStr, ok := request["provider_id"].(string)
	if ok {
		payload.ProviderID, _ = primitive.ObjectIDFromHex(supervisorIdStr)
	}

	adminIdStr, ok := request["admin_id"].(string)
	if ok {
		payload.AdminID, _ = primitive.ObjectIDFromHex(adminIdStr)
	}

	supplyIdStr, ok := request["supply_id"].(string)
	if ok {
		payload.SupplyID, _ = primitive.ObjectIDFromHex(supplyIdStr)
	}

	fromIdStr, ok := request["from_id"].(string)
	if ok {
		payload.FromID, _ = primitive.ObjectIDFromHex(fromIdStr)
	}

	destinationIdStr, ok := request["destination_id"].(string)
	if ok {
		payload.DestinationID, _ = primitive.ObjectIDFromHex(destinationIdStr)
	}

	reason, ok := request["reason"].(string)
	if ok {
		payload.Reason = reason
	}

	status, ok := request["status"].(string)
	if ok {
		/*
		 "pending", "approved", "dismissed"
		*/
		payload.Status = _entities.NotificationStatus(status)
	}

	payload.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	supervisor, err := _mongo.FetchSupervisorById(payload.ProviderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = _mongo.UpsertNewApproval(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hub := _websockets.GetSocketHub()
	hub.SendToUser(payload.AdminID, map[string]interface{}{
		"message": fmt.Sprintf("Your supervisor has asked for order-approval. It is in %s status right now. Please take action.", status),
		"success": true,
	})

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("Supervisor %s has asked for order approval.", supervisor.Name),
		"success": true,
	})
}

func UpdateFinalOrderApproval(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request", http.StatusBadRequest);
		return;
	}

	currentStatus := r.URL.Query().Get("current_status")
	orderIdStr := r.URL.Query().Get("order_id")

	adminIdStr := r.URL.Query().Get("admin_id")
	adminID, _ := primitive.ObjectIDFromHex(adminIdStr)

	toWarehouseIdStr := r.URL.Query().Get("to_warehouse_id");
	siteIdStr := r.URL.Query().Get("site_id");
	
	materialName := r.URL.Query().Get("material_name");
	approvalIdStr := r.URL.Query().Get("approval_id");


	var orderID primitive.ObjectID
	var toWarehouseID primitive.ObjectID;
	var siteID primitive.ObjectID;
	var approvalID primitive.ObjectID;

	if len(orderIdStr) > 0 {
		orderID, _ = primitive.ObjectIDFromHex(orderIdStr)
	}

	if len(toWarehouseIdStr) > 0 {
		toWarehouseID,_ = primitive.ObjectIDFromHex(toWarehouseIdStr);
	}

	if len(siteIdStr) > 0 {
		siteID,_ = primitive.ObjectIDFromHex(siteIdStr);
	} else {
		siteID = primitive.NilObjectID
	}

	if len(approvalIdStr) > 0 {
		approvalID ,_ = primitive.ObjectIDFromHex(approvalIdStr);
	}


	err := _mongo.UpdateFinalOrderStatus(currentStatus, materialName, orderID, toWarehouseID, siteID, approvalID)

	if err != nil {
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}


	hub := _websockets.GetSocketHub()
		hub.SendToUser(adminID, map[string]interface{}{
			"message": fmt.Sprintf("Order ID %s status is updated to %s.Congratulations.", orderIdStr, currentStatus),
		})
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("The status of OrderID-%s is %s ", orderIdStr, currentStatus),
		})
		return

}

func UpdateOrderApproval(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request", http.StatusBadRequest)
		return
	}

	currentStatus := r.URL.Query().Get("current_status")
	orderIdStr := r.URL.Query().Get("order_id")
	adminIdStr := r.URL.Query().Get("admin_id")
	adminID, _ := primitive.ObjectIDFromHex(adminIdStr)

	var orderID primitive.ObjectID

	if len(orderIdStr) > 0 {
		orderID, _ = primitive.ObjectIDFromHex(orderIdStr)
	}


	if currentStatus == string(_entities.OutForDelivery) {
		// out for delivery.
		err := _mongo.UpdateOrder(orderID, currentStatus)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hub := _websockets.GetSocketHub()
		hub.SendToUser(adminID, map[string]interface{}{
			"message": fmt.Sprintf("Order ID %s status is updated to %s. Kindly look into it.", orderIdStr, currentStatus),
		})

		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("The status of OrderID-%s is %s ", orderIdStr, currentStatus),
		})
		return
	}
}
