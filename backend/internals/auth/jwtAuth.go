package auth

import (
	_envConfig "atheena/internals/config"
	_mongo "atheena/internals/database/mongoV2"
	entities "atheena/internals/entities"
	util "atheena/internals/util"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
)

func GenerateRefreshToken() string {
	log.Println("Generating Refresh Token....");
	return util.GenerateRandomUUID()
}



// access_token, refresh_token, error, http_status (Generates new tokens and inserts new user)
func GenerateNewAccessAndRefreshTokens(user entities.User) (*string, *string, error, int){

	access_token , err := GenerateAccessToken(user.Email, user.Role);

	if err != nil {
		log.Println("Something went wrong, when access_token was generated !");
		return nil, nil, err, http.StatusInternalServerError;
	}


	refresh_token := GenerateRefreshToken()


	// Set 1 day as expiry of refresh token...
	expiry, _ :=  time.Parse( time.RFC3339, time.Now().Add(24 * time.Hour).Format(time.RFC3339));

	created_time,_ := util.GenerateCreateDateTime()

	// Here we will have expiry of refresh token
	authToken := entities.AuthToken{
		ID: util.GenerateObjectID(),
		User_Id: user.ID,
		Email: user.Email,
		Role: user.Role,
		Refresh_Token: refresh_token,
		Expiry_Time: expiry,
		Created_At: created_time,
	}

	err = _mongo.InsertAuthToken(authToken);
	if (err != nil) {
		log.Println("Something went wrong, while inserting new token !");
		return nil, nil, err, http.StatusInternalServerError;
	}

	return &access_token, &refresh_token, nil, http.StatusOK;
}

func RenewAccessToken(refreshToken string) (*string, error, int) {

	authToken, err := _mongo.GetTokenByRefreshToken(refreshToken);

	if err != nil {
		log.Println("Something went wrong  while renewing access token");
		return nil, err, http.StatusInternalServerError;
	}

	log.Println("The time.Now() is ", util.FormatDateTime(time.Now()))
	log.Println("The expiry time of token is ", util.FormatDateTime(authToken.Expiry_Time))

	if (time.Now().After(authToken.Expiry_Time)) {
		return nil, fmt.Errorf("expired token"), http.StatusForbidden;
	}

	newAccessToken, err := GenerateAccessToken(authToken.Email, authToken.Role)
	if err != nil {
		log.Println("Something went wrong while getting new access tokens.");
		log.Println(err.Error());
		return nil, err, http.StatusInternalServerError
	}

	return &newAccessToken, nil, http.StatusOK;
}

func GenerateAccessToken(email, role string) (string ,error) {
	log.Println("Generating access tokens...")

	claims := jwt.MapClaims{
		util.JWT_USER_EMAIL:email,
		util.JWT_USER_ROLE:role,
		// Expiry time need to 7 mins, but for now setting it to 10 mins.
		util.JWT_USER_EXPIRATION: time.Now().Add(10 * time.Minute).Unix(),
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envConfig := _envConfig.LoadEnvConfig();

	return access_token.SignedString([]byte(envConfig.JWT_Secret));
}

