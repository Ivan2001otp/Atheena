package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	_mongo "atheena/internals/database/mongoV2"
	_entities "atheena/internals/entities"
)


func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method", http.StatusBadRequest);
		return;
	}

	var requestPayload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestPayload);

	if err != nil {
		log.Println("could not parse request body to delete supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	var adminId, superVisorId primitive.ObjectID
	adminIdStr, ok := requestPayload["admin_id"].(string)
	if (ok) {
		adminId,_ = primitive.ObjectIDFromHex(adminIdStr)
	}

	supervisorIdStr , ok :=  requestPayload["id"].(string)
	if (ok) {
		superVisorId,_ = primitive.ObjectIDFromHex(supervisorIdStr)
	}


	err = _mongo.DeleteSupervisor(superVisorId, adminId)
	if err != nil {
		log.Println("Something went wrong while deleting supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	response := map[string]interface{}{
		"success":true,
		"message":fmt.Sprintf("%s - supervisor is removed successfully.", superVisorId.Hex()),
	}

	json.NewEncoder(w).Encode(response);
}

func AddOrUpdateSupervisor(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method", http.StatusBadRequest);
		return;
	}

	var requestPayload map[string]interface{}
	var supervisor _entities.Supervisor

	err := json.NewDecoder(r.Body).Decode(&requestPayload);

	if err != nil {
		log.Println("could not parse request body to create supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	supervisor.ID = primitive.NewObjectID()

	adminIdStr , ok := requestPayload["admin_id"].(string)
	if (ok){
	  supervisor.AdminID ,_ = primitive.ObjectIDFromHex(adminIdStr)
	}

	nameStr, ok := requestPayload["name"].(string)
	if (ok) {
		supervisor.Name = nameStr
	}

	emailStr, ok := requestPayload["email"].(string)
	if (ok) {
		supervisor.Email = emailStr;
	}

	phoneNumStr, ok := requestPayload["phone_number"].(string)
	if (ok) {
		supervisor.PhoneNumber = phoneNumStr;
	}

	roleStr, ok := requestPayload["role"].(string)
	if (ok) {
		supervisor.Role = roleStr;
	}

	supervisor.Created_At,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

	err=_mongo.UpsertNewSupervisor(supervisor)
	if err != nil {
		log.Println("Something wrong happened while upserting supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return;
	}


	response := map[string]interface{}{
		"success":true,
		"message":fmt.Sprintf("%s - supervisor added successfully.", supervisor.Name),
	}

	_ = json.NewEncoder(w).Encode(response)

}