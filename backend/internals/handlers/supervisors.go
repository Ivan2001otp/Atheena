package handlers


import (
	"log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

	_entities "atheena/internals/entities"
	_mongo "atheena/internals/database/mongoV2"
)


func DeleteSupervisor(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method", http.StatusBadRequest);
		return;
	}

	var supervisor _entities.Supervisor
	err := json.NewDecoder(r.Body).Decode(&supervisor);

	if err != nil {
		log.Println("could not parse request body to delete supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	err = _mongo.DeleteSupervisor(supervisor)
	if err != nil {
		log.Println("Something went wrong while deleting supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	response := map[string]interface{}{
		"success":false,
		"message":fmt.Sprintf("%s - supervisor is removed successfully."),
	}

	json.NewEncoder(w).Encode(response);
}

func AddOrUpdateSupervisor(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method", http.StatusBadRequest);
		return;
	}

	var supervisor _entities.Supervisor
	err := json.NewDecoder(r.Body).Decode(&supervisor);

	if err != nil {
		log.Println("could not parse request body to create supervisor");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	supervisor.ID = primitive.NewObjectID()
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