package handlers

import (
	_entities "atheena/internals/entities"
	"encoding/json"
	"log"
	"net/http"
	"time"
	_websockets "atheena/internals/database/websockets"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AskForApproval(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		log.Println("supposed to be POST request");
		http.Error(w, "Supposed to be POST request.", http.StatusBadRequest);
		return;
	}

	var request map[string] interface{}

	err := json.NewDecoder(r.Body).Decode(&request);
	if err != nil {
		log.Println(err.Error());
		log.Println("Failed to parse the request body.");
		return;
	}

	var payload _entities.ApprovalTypeNotification;
	payload.ApprovalID = primitive.NewObjectID()

	supervisorIdStr, ok := request["provider_id"].(string)
	if (ok) {
		payload.ProviderID,_ = primitive.ObjectIDFromHex(supervisorIdStr);
	}

	adminIdStr, ok := request["admin_id"].(string)
	if (ok) {
		payload.AdminID,_ = primitive.ObjectIDFromHex(adminIdStr);
	}

	supplyIdStr, ok := request["supply_id"].(string)
	if (ok) {
		payload.SupplyID,_ = primitive.ObjectIDFromHex(supplyIdStr);
	}

	orderIdStr , ok := request["order_id"].(string)
	if (ok) {
		payload.OrderID,_ = primitive.ObjectIDFromHex(orderIdStr);
	}

	fromIdStr, ok := request["from_id"].(string)
	if (ok) {
		payload.FromID,_ = primitive.ObjectIDFromHex(fromIdStr);
	}

	destinationIdStr, ok := request["destination_id"].(string)
	if (ok) {
		payload.DestinationID,_ = primitive.ObjectIDFromHex(destinationIdStr);
	}

	reason, ok := request["reason"].(string)
	if (ok) {
		payload.Reason = reason;
	}

	status, ok := request["status"].(string)
	if (ok) {
		/*
			 "pending", "approved", "dismissed"
		*/
		payload.Status = _entities.NotificationStatus(status);
	}

	payload.CreatedAt,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

    //log.Println(payload)

	hub:= _websockets.GetSocketHub()
	hub.SendToUser(payload.AdminID, &payload);
}