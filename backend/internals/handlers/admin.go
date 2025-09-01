package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	_mongo "atheena/internals/database/mongoV2"
	_entities "atheena/internals/entities"
	// "atheena/internals/util"
)

// supervisor related CRUD.
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



// Adding inventory item to a particular warehouse id.
func AddInventoryItem(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST request", http.StatusBadRequest);
		return;
	}

	var inventoryItem 	_entities.InventoryItem;
	err := json.NewDecoder(r.Body).Decode(&inventoryItem)
	
	if err != nil {
		log.Println(err.Error());
		log.Println("Something went wrong while parsing the rquest body..")
	}	
}



// Fetch the constructions sites of the given admin id
func FetchConstructionSitebyAdminId(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "Only GET request", http.StatusBadRequest);
		return;
	}

	params := mux.Vars(r);
	adminIdStr := params["admin_id"]
	log.Println("The admin id is ", adminIdStr);

	var constructionSites []_entities.Site;
	adminId,_ := primitive.ObjectIDFromHex(adminIdStr);

	constructionSites, err := _mongo.FetchSitesbyAdminId(adminId);

	if err != nil {
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}


	log.Println("✅ Successfully fetched constructions sites for admin id", adminIdStr);
	_ = json.NewEncoder(w).Encode(constructionSites);
}

// Construction Site related CRUD.
func AddConstructionSite(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST request", http.StatusBadRequest);
		return;
	}

	var constructionSite _entities.Site;
	var requestPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestPayload);
	if err != nil {
		log.Println("Something went wrong while parsing request body...");
		log.Println(err.Error());
		return;
	}

	nameStr, ok := requestPayload["name"].(string)
	if (ok) {
		constructionSite.Name = nameStr;
	}


	addressStr, ok := requestPayload["address"].(string)
	if (ok) {
		constructionSite.Address = addressStr;
	}

	stateStr, ok := requestPayload["state"].(string)
	if (ok) {
		constructionSite.State = stateStr;
	}

	countryStr, ok := requestPayload["country"].(string)
	if (ok) {
		constructionSite.Country = countryStr;
	}

	adminIdStr, ok := requestPayload["user_id"].(string)
	if (ok) {
		fmt.Println("Admin id str is ", adminIdStr);
		constructionSite.AdminId,_ = primitive.ObjectIDFromHex(adminIdStr)
		fmt.Println("Admin id object id is ", constructionSite.AdminId);
	}


	constructionSite.ID = primitive.NewObjectID()
	constructionSite.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

	err = _mongo.UpsertNewConstructionSiteesByAdmin(constructionSite)
	if err != nil {
		log.Println("Failed to upsert the construction site record");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	log.Println("Successfully added/updated construction site record");
	response := map[string]interface{}{
		"success":true,
		"message":"Successfully updated/added construction site record.",
	}

	_ = json.NewEncoder(w).Encode(response);
}



// adding warehouses.
func AddNewWarehouse(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST request", http.StatusBadRequest);
		return;
	}


	var requestPayload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestPayload);
	if err != nil {
		log.Println("Could not parse warehouse request body.");
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	var warehouse _entities.WareHouse
	userIdStr,ok := requestPayload["user_id"].(string)
	log.Println("The real user_id value from frontend : ",userIdStr);
	if (ok) {
		warehouse.User_Id,_ = primitive.ObjectIDFromHex(userIdStr)
	}
	
	nameStr, ok := requestPayload["name"].(string)
	if (ok) {
		warehouse.Name = nameStr;
	}

	pinStr, ok := requestPayload["pin"].(string)
	if (ok) {
		warehouse.Pin = pinStr
	}

	addressStr, ok := requestPayload["address"].(string)
	if (ok) {
		warehouse.Address = addressStr
	}

	stateStr, ok := requestPayload["state"].(string)
	if (ok) {
		warehouse.State = stateStr
	}

	countryStr, ok := requestPayload["country"].(string)
	if (ok) {
		warehouse.Country = countryStr
	}

	idStr, ok := requestPayload["id"].(string);
	if (ok) {

		if (len(idStr)==0){
			warehouse.ID = primitive.NewObjectID();
		} else {
			warehouse.ID,_ = primitive.ObjectIDFromHex(idStr);
		}
	}

	// warehouse.ID = primitive.NewObjectID()
	warehouse.Created_At ,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
	err = _mongo.AddWarehouseByUser(warehouse);
	
	if err != nil {
		log.Println("Something went wrong while adding new warehouse..");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	response := map[string]interface{}{
		"success":true,
		"message":fmt.Sprintf("Successfully added %s warehouse", warehouse.Name),
	}

	_ = json.NewEncoder(w).Encode(response);
}

// get all warehouses for given admin_id
func GetAllWarehouseByAdminId(w http.ResponseWriter, r *http.Request) {
	
	if (r.Method != http.MethodGet) {
		http.Error(w, "Only GET request", http.StatusBadRequest);
		return;
	}

	params := mux.Vars(r);
	adminIdStr := params["admin_id"];

	var wareHouseList []_entities.WareHouse;
	adminId,_ := primitive.ObjectIDFromHex(adminIdStr);

	wareHouseList, err := _mongo.FetchWarehouseById(adminId);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	log.Println("✅ Successfully fetched warehouses for admin id", adminIdStr);
	_ = json.NewEncoder(w).Encode(wareHouseList);
}


// This function need to be revisited again.
func RemoveWarehouse(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodDelete) {
		http.Error(w, "Supposed to be DELETE", http.StatusBadRequest);
		return;
	}

	warehouseIdStr := r.URL.Query().Get("warehouse_id");

	warehouseId,_ := primitive.ObjectIDFromHex(warehouseIdStr)
	err := _mongo.DeleteWarehouseById(warehouseId)

	if err != nil {
		log.Println("Something went wrong while deleting warehouse.");
		log.Println(err);
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	response := map[string]interface{}{
		"success":true,
		"message":fmt.Sprintf("Warehouse-%s is successfully removed.", warehouseIdStr),
	}

	_ = json.NewEncoder(w).Encode(response);
}

