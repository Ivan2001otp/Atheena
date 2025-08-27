package handlers

import (
	
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	_jwtAuth "atheena/internals/auth"
	_mongoRepo "atheena/internals/database/mongoV2"
	"atheena/internals/entities"
)

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Here we delete the user entry by Id
	// The api will recieve this Id as "string" format
	// We have to convert it to primitive.ObjectID
	objectIDstr := r.URL.Query().Get("object-id")
	objectID,err := primitive.ObjectIDFromHex(objectIDstr)
	
	if err != nil {
		log.Println("Something went wrong while converting string objective-id to primitive objective-id.");
		return;
	}

	err = _mongoRepo.DeleteUserById(objectID)
	if err != nil {
		log.Println("Something went wrong while deleting user by object-id");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	response := map[string]interface{}{
		"success":"true",
		"message":"logout successfully",
	}

	_ = json.NewEncoder(w).Encode(response);
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email");
	role := r.URL.Query().Get("role");

	// soft delete the user's token entry from auth-tokens.
	err := _mongoRepo.DeleteLoggedOutRefreshToken(email, role);
	if err != nil {
		log.Println("LogoutHandler()");
		log.Println(err.Error());
		http.Error(w, "Something went wrong during logout !", http.StatusInternalServerError);
		return;
	}


	response := map[string]interface{}{
		"success":true,
		"message":"logout successfully",
	}

	_= json.NewEncoder(w).Encode(response)

}

// Login with email and Password
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed);
		return;
	}


	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user);

	if err != nil {
		log.Println("Failed to parse request body");
		log.Println(err.Error());
		http.Error(w, "Failed to parse request body", http.StatusInternalServerError);
		
		return;
	}

	// Check whether the email exists or not
	fetchedUser, err :=  _mongoRepo.EmailExists(user.Email)
	
	if (fetchedUser == nil && err == nil) {
		// email does not exists
		log.Println("Email does not exists !");
		http.Error(w, "Email does not exists", http.StatusUnauthorized);
		return;
	}

	if err != nil {
		log.Println("Something went wrong while checking email exists or not during login flow.");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}


	// Check whether the password is correct or not...
	err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(user.Password));
	if err != nil {
		
		if (err == bcrypt.ErrMismatchedHashAndPassword) {
			
			log.Println("password did not match");
			http.Error(w, "Invalid Password", http.StatusUnauthorized);
			return;
		}
		log.Println("Something went wrong while comparing hash and password !");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}


	access_token, refresh_token, err , status_code := _jwtAuth.GenerateNewAccessAndRefreshTokens(*fetchedUser);
	if err != nil {
		log.Println("The status code is ", status_code);
		log.Println("Something went wrong while generating tokens during login.");
		http.Error(w, "Something went wrong.", status_code);
		return;
	}

	response := map[string]interface{}{
		"message":"success",
		"access_token":access_token,
		"refresh_token":refresh_token,
	}


	json.NewEncoder(w).Encode(response);
}


// Register with Email and Password.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed);
		return;
	}

	var user entities.User;
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Could not parse the user request body");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	// Check whether the email exists or not
	fetchedUser, err :=  _mongoRepo.EmailExists(user.Email)
	if err != nil {
		log.Println("Something went wrong while checking user email during register !");
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}

	if (fetchedUser != nil) {
		log.Println("Email already exists");
		response := map[string]interface{}{
			"message":fmt.Sprintf("Email %s already exists.", user.Email),
		}

		_ = json.NewEncoder(w).Encode(response);
		return;
	}

	log.Println("email : ", user.Email);
	log.Println("role : ", user.Role);
	log.Println("name : ", user.Name);

	// hash the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost);
	if (err != nil) {
		log.Println("failed to hash the given password during register user");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return ;
	}

	user.ID = primitive.NewObjectID();
	user.Password = string(hashedPassword);
	user.Visited_Time,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339));
	log.Println("Printing the user model");
	log.Println(user);

	// add new user to db....
	err = _mongoRepo.InsertNewUser(user)
	if err != nil {
		log.Println("Could not insert user !");
		log.Println(err.Error());
		http.Error(w, err.Error(), http.StatusInternalServerError);
		return;
	}


	// get new tokens
	accessToken, refreshToken, err, status := _jwtAuth.GenerateNewAccessAndRefreshTokens(user);
	if (err != nil) {
		log.Println("Something went wrong while generating access_token and refresh_token");
		http.Error(w, err.Error(), status);
		return;
	}


	w.WriteHeader(http.StatusOK);
	response := map[string]interface{}{
		"access_token" : accessToken,
		"refresh_token": refreshToken,
		"message":"success",
	}


	_ = json.NewEncoder(w).Encode(response);
}