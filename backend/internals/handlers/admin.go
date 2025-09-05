package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	_mongo "atheena/internals/database/mongoV2"
	_entities "atheena/internals/entities"
	// "atheena/internals/util"
)

// logs related CRUD
func FetchAllLogs(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "Only GET request is allowed", http.StatusBadRequest);
		return;
	}

	admin_id := r.URL.Query().Get("admin_id");
	if (len(admin_id) == 0) {
		http.Error(w, "admin_id is missing in params.", http.StatusBadRequest);
		return;
	}

	adminId,_ := primitive.ObjectIDFromHex(admin_id);
	err,response := _mongo.FetchLogs(adminId);

	if err != nil {
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	log.Println("✅ Logs fetched succesfully");
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success":true,
		"message":"Fetched logs successfully",
		"data":response,
	});
}


func FetchOrders(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "only GET request valid.", http.StatusBadRequest);
		return;
	}

	adminIdStr := r.URL.Query().Get("admin_id");
	adminId, _ := primitive.ObjectIDFromHex(adminIdStr);

	err , result := _mongo.FetchOrders(adminId);
	if err != nil {
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success":true,
		"message":"Fetched Order successfully",
		"data":result,
	});
}


// approval related CRUD
func FetchAllApprovals(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "only GET request valid.", http.StatusBadRequest);
		return;
	}

	adminIdStr := r.URL.Query().Get("admin_id");
	adminId ,_ := primitive.ObjectIDFromHex(adminIdStr);

	err, approvalList := _mongo.FetchAllApprovals(adminId);

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;	
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"data":approvalList,
		"message":"Fetched approvals successfully.",
		"success":true,
	});

}

// supervisor related CRUD.
func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method", http.StatusBadRequest)
		return
	}

	var requestPayload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestPayload)

	if err != nil {
		log.Println("❌could not parse request body to delete supervisor")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var adminId, superVisorId primitive.ObjectID
	adminIdStr, ok := requestPayload["admin_id"].(string)
	if ok {
		adminId, _ = primitive.ObjectIDFromHex(adminIdStr)
	}

	supervisorIdStr, ok := requestPayload["id"].(string)
	if ok {
		superVisorId, _ = primitive.ObjectIDFromHex(supervisorIdStr)
	}

	err = _mongo.DeleteSupervisor(superVisorId, adminId)
	if err != nil {
		log.Println("❌Something went wrong while deleting supervisor")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("%s - supervisor is removed successfully.", superVisorId.Hex()),
	}

	json.NewEncoder(w).Encode(response)
}

func AddOrUpdateSupervisor(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method", http.StatusBadRequest)
		return
	}

	var requestPayload map[string]interface{}
	var supervisor _entities.Supervisor

	err := json.NewDecoder(r.Body).Decode(&requestPayload)

	if err != nil {
		log.Println("could not parse request body to create supervisor")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(requestPayload);
	supervisrIdStr, ok := requestPayload["id"].(string);
	log.Println("supervisorIdStr is ", supervisrIdStr);
	
	if ok {
		if (len(supervisrIdStr) > 0) {
			log.Println("supervisor id exists.");
			supervisor.ID, _ = primitive.ObjectIDFromHex(supervisrIdStr);
		} 
	} else {
		log.Println("supervisor id is created.");
		supervisor.ID = primitive.NewObjectID();
	}

	adminIdStr, ok := requestPayload["admin_id"].(string)
	if ok {
		supervisor.AdminID, _ = primitive.ObjectIDFromHex(adminIdStr)
	}

	nameStr, ok := requestPayload["name"].(string)
	if ok {
		supervisor.Name = nameStr
	}

	emailStr, ok := requestPayload["email"].(string)
	if ok {
		supervisor.Email = emailStr
	}

	phoneNumStr, ok := requestPayload["phone_number"].(string)
	if ok {
		supervisor.PhoneNumber = phoneNumStr
	}

	roleStr, ok := requestPayload["role"].(string)
	if ok {
		supervisor.Role = roleStr
	}

	
	err = _mongo.UpsertNewSupervisor(supervisor)
	if err != nil {
		log.Println("❌Something wrong happened while upserting supervisor")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("%s - supervisor added successfully.", supervisor.Name),
	}

	_ = json.NewEncoder(w).Encode(response)
}

func FetchAllSupervisor(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "Only GET request", http.StatusBadRequest);
		return;
	}

	adminIdStr := r.URL.Query().Get("admin_id");
	adminId,_ := primitive.ObjectIDFromHex(adminIdStr)
	supervisors, err := _mongo.FetchAllSupervisorByAdminId(adminId);

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success":true,
		"message":"Fetched supervisors successfully",
		"data":supervisors,
	});
}

// Adding inventory item to a particular warehouse id.
func AddInventoryItem(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request", http.StatusBadRequest)
		return
	}

	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println(err.Error())
		log.Println("❌Something went wrong while parsing the rquest body..")
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	log.Println("inventory payload")
	log.Println(payload)
	var inventoryItem _entities.InventoryItem
	inventoryItem.ID = primitive.NewObjectID()

	warehouseIdStr, ok := payload["warehouse_id"].(string)
	if ok {
		inventoryItem.Warehouse_Id, _ = primitive.ObjectIDFromHex(warehouseIdStr)
	}

	name, ok := payload["name"].(string)
	if ok {
		inventoryItem.Name = name
	}

	quantity, ok := payload["quantity"].(float64)

	if ok {
		inventoryItem.Quantity = quantity
	}

	unit, ok := payload["unit"].(string)
	if ok {
		inventoryItem.Unit = unit
	}

	reason, ok := payload["reason"].(string)
	if ok {
		inventoryItem.Reason = reason
	}

	inventoryItem.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	inventoryItem.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))



	err = _mongo.AddNewInventoryItems(inventoryItem)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("✅ Upserted inventory Item successfully")

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Upserted %s inventory successfully", inventoryItem.Name),
	}
	_ = json.NewEncoder(w).Encode(response)
}

func FetchInventoryByWarehouse(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodGet) {
		http.Error(w, "Only GET request is available.", http.StatusBadRequest);
		return;
	}

	warehouseIdStr := r.URL.Query().Get("warehouse_id");
	if ( len(warehouseIdStr) == 0) {
		log.Println("warehouseId is empty");
		http.Error(w, "No warehouseId given as parameter to the API.", http.StatusBadRequest);
		return;
	}

	warehouseId,_ := primitive.ObjectIDFromHex(warehouseIdStr);

	var inventoryList []_entities.InventoryItem;

	inventoryList, err := _mongo.FetchInventoryByWarehouseId(warehouseId);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success":true,
		"message":"Fetched inventories successfully",
		"data":inventoryList,
	})

}

// Fetch the constructions sites of the given admin id
func FetchConstructionSitebyAdminId(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET request", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	adminIdStr := params["admin_id"]
	log.Println("The admin id is ", adminIdStr)

	var constructionSites []_entities.Site
	adminId, _ := primitive.ObjectIDFromHex(adminIdStr)

	constructionSites, err := _mongo.FetchSitesbyAdminId(adminId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("✅ Successfully fetched constructions sites for admin id", adminIdStr)
	_ = json.NewEncoder(w).Encode(constructionSites)
}

// Construction Site related CRUD.
func AddConstructionSite(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request", http.StatusBadRequest)
		return
	}

	var constructionSite _entities.Site
	var requestPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println("❌Something went wrong while parsing request body...")
		log.Println(err.Error())
		return
	}

	nameStr, ok := requestPayload["name"].(string)
	if ok {
		constructionSite.Name = nameStr
	}

	addressStr, ok := requestPayload["address"].(string)
	if ok {
		constructionSite.Address = addressStr
	}

	stateStr, ok := requestPayload["state"].(string)
	if ok {
		constructionSite.State = stateStr
	}

	countryStr, ok := requestPayload["country"].(string)
	if ok {
		constructionSite.Country = countryStr
	}

	adminIdStr, ok := requestPayload["user_id"].(string)
	if ok {
		fmt.Println("Admin id str is ", adminIdStr)
		constructionSite.AdminId, _ = primitive.ObjectIDFromHex(adminIdStr)
		fmt.Println("Admin id object id is ", constructionSite.AdminId)
	}

	constructionSite.ID = primitive.NewObjectID()
	constructionSite.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	err = _mongo.UpsertNewConstructionSiteesByAdmin(constructionSite)
	if err != nil {
		log.Println("❌Failed to upsert the construction site record")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Successfully added/updated construction site record")
	response := map[string]interface{}{
		"success": true,
		"message": "Successfully updated/added construction site record.",
	}

	_ = json.NewEncoder(w).Encode(response)
}

// adding warehouses.
func AddNewWarehouse(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request", http.StatusBadRequest)
		return
	}

	var requestPayload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestPayload)
	if err != nil {
		log.Println("❌Could not parse warehouse request body.")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var warehouse _entities.WareHouse
	userIdStr, ok := requestPayload["user_id"].(string)
	log.Println("The real user_id value from frontend : ", userIdStr)
	if ok {
		warehouse.User_Id, _ = primitive.ObjectIDFromHex(userIdStr)
	}

	nameStr, ok := requestPayload["name"].(string)
	if ok {
		warehouse.Name = nameStr
	}

	pinStr, ok := requestPayload["pin"].(string)
	if ok {
		warehouse.Pin = pinStr
	}

	addressStr, ok := requestPayload["address"].(string)
	if ok {
		warehouse.Address = addressStr
	}

	stateStr, ok := requestPayload["state"].(string)
	if ok {
		warehouse.State = stateStr
	}

	countryStr, ok := requestPayload["country"].(string)
	if ok {
		warehouse.Country = countryStr
	}

	idStr, ok := requestPayload["id"].(string)
	if ok {

		if len(idStr) == 0 {
			warehouse.ID = primitive.NewObjectID()
		} else {
			warehouse.ID, _ = primitive.ObjectIDFromHex(idStr)
		}
	}

	// warehouse.ID = primitive.NewObjectID()
	warehouse.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	err = _mongo.AddWarehouseByUser(warehouse)

	if err != nil {
		log.Println("❌Something went wrong while adding new warehouse..")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully added %s warehouse", warehouse.Name),
	}

	_ = json.NewEncoder(w).Encode(response)
}

// get all warehouses for given admin_id
func GetAllWarehouseByAdminId(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET request", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	adminIdStr := params["admin_id"]

	var wareHouseList []_entities.WareHouse
	adminId, _ := primitive.ObjectIDFromHex(adminIdStr)

	wareHouseList, err := _mongo.FetchWarehouseById(adminId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("✅ Successfully fetched warehouses for admin id", adminIdStr)
	_ = json.NewEncoder(w).Encode(wareHouseList)
}

// This function need to be revisited again.
func RemoveWarehouse(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Supposed to be DELETE", http.StatusBadRequest)
		return
	}

	warehouseIdStr := r.URL.Query().Get("warehouse_id")

	warehouseId, _ := primitive.ObjectIDFromHex(warehouseIdStr)
	err := _mongo.DeleteWarehouseById(warehouseId)

	if err != nil {
		log.Println("❌Something went wrong while deleting warehouse.")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Warehouse-%s is successfully removed.", warehouseIdStr),
	}

	_ = json.NewEncoder(w).Encode(response)
}


// Order Tracking handlers 
func InteractOrderApproval(w http.ResponseWriter, r *http.Request) {
	// It will use admin_id, approval_id, updated_status

	if (r.Method != http.MethodGet) {
		log.Println("Only GET request allowed.")
		http.Error(w, "Only GET request allowed.", http.StatusBadRequest);
		return;
	}

	admin_id_str := r.URL.Query().Get("admin_id");
	adminId,_ := primitive.ObjectIDFromHex(admin_id_str);

	approval_id_str :=  r.URL.Query().Get("approval_id");
	approvalId,_ := primitive.ObjectIDFromHex(approval_id_str);

	updated_status := r.URL.Query().Get("updated_status");
	

	supervisor_id_str := r.URL.Query().Get("supervisor_id");
	supervisorId, _ := primitive.ObjectIDFromHex(supervisor_id_str);

	from_warehouse_id_str := r.URL.Query().Get("from_warehouse_id");
	fromWareHouseID, _ := primitive.ObjectIDFromHex(from_warehouse_id_str); 

	to_warehouse_id_str := r.URL.Query().Get("to_warehouse_id");
	toWareHouseID,_ := primitive.ObjectIDFromHex(to_warehouse_id_str);

	site_id_str := r.URL.Query().Get("site_id");
	var siteID primitive.ObjectID;
	if site_id_str == "" {
		siteID = primitive.NilObjectID;
	}else {siteID,_ = primitive.ObjectIDFromHex(site_id_str);}

	log_type_str := r.URL.Query().Get("log_type");
	updated_time,_ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));

	quantity := r.URL.Query().Get("quantity")
	quantityF64, _ := strconv.ParseFloat(quantity, 64);

	inventory_id_str := r.URL.Query().Get("inventory_id");
	inventoryId, _ := primitive.ObjectIDFromHex(inventory_id_str);

	material_name := r.URL.Query().Get("material");
	unit := r.URL.Query().Get("unit");

	var logRecord _entities.CustomLog
	logRecord.ID = primitive.NewObjectID()
	logRecord.AdminId = adminId
	logRecord.Site_Id = siteID;
	logRecord.SupplyID = inventoryId;
	logRecord.Supervisor_ID = supervisorId;
	logRecord.From_Warehouse_Id = fromWareHouseID;
	logRecord.To_Warehouse_Id = toWareHouseID;
	logRecord.Log_Type = _entities.LogType(log_type_str);
	logRecord.Updated = updated_time;

	
	// update the status of approval-record
	err := _mongo.UpdateApprovalStatus(approvalId, updated_status);
	if err != nil {
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	// add / update the log-record.
	err = _mongo.UpsertLog(logRecord);
	if err != nil {
		log.Println(err.Error());
		log.Println("Something went wrong while upserting the log record.");
		http.Error(w, "Something went wrong while upserting the log record.", http.StatusInternalServerError);
		return;
	}
	
	if (updated_status != "approved") {
		w.WriteHeader(http.StatusOK);
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message":fmt.Sprintf("The request for material transit is %s", updated_status),
			"success":true,
		});

		return;
	}

	// update the inventory of source.
	err = _mongo.UpdateInventoryQuantity(fromWareHouseID, inventoryId, quantityF64);
	if err != nil {
		log.Println(err.Error());
		log.Println("Something went wrong while updating inventory quantity.");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	// create the order when status is approved.
	// log_type_str and order_type_str are same.
	err = _mongo.CreateOrderRecord(logRecord.ID, quantityF64, log_type_str, material_name, unit);
	if (err != nil) {
		log.Println(err.Error());
		log.Println("❌Something went wrong while creating order.");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":"Order created successfully",
		"success":true,
	});
	log.Println("✅ Order created successfully.");
}
