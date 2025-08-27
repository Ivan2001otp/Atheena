package auth

import (
	"github.com/dgrijalva/jwt-go"
	util "atheena/internals/util"
	_envConfig "atheena/internals/config"
	"log"
	"time"
)

func GenerateRefreshToken() string {
	log.Println("Generating Refresh Token....");
	return util.GenerateRandomUUID()
}

func GenerateAccessToken(email, role string) (string ,error) {
	log.Println("Generating access tokens...")

	claims := jwt.MapClaims{
		util.JWT_USER_EMAIL:email,
		util.JWT_USER_ROLE:role,
		// Expiry time need to 7 mins, but for now setting it to 15 mins.
		util.JWT_USER_EXPIRATION: time.Now().Add(15 * time.Minute).Unix(),
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envConfig := _envConfig.LoadEnvConfig();

	return access_token.SignedString([]byte(envConfig.JWT_Secret));

}

