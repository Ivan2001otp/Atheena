package middleware

import (
	envConfig "atheena/internals/config"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/time/rate"
)


func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization");
			

			if (! strings.HasPrefix(authHeader, "Bearer ")) {
				log.Println("The token is malformed in header.")
				http.Error(w, "Missing or Malformed token", http.StatusUnauthorized);
				return;
			}


			tokenStr := strings.TrimPrefix(authHeader, "Bearer ");
			log.Println("Token sent in header : ", tokenStr);
			secretKey := envConfig.LoadEnvConfig().JWT_Secret;

			// Try to unmarshal the token to understand its authenticity.
			token,err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil;
			})


			if (token == nil) {
				log.Println("The unmarshalled token is NULL.");
				http.Error(w, "The unmarshalled token is NULL.", http.StatusInternalServerError);
				return;
			}

			if (err != nil) {
				log.Println("Something wrong happend Immanuel");
				log.Println(err.Error());
				// checking whether the token is expired or not.
				if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
					http.Error(w, "Access Token expired" , http.StatusUnauthorized);
				} else {
					http.Error(w, "Invalid Token . Please contact the team", http.StatusUnauthorized);
				}

				return;
			}


			if (!token.Valid) {
				http.Error(w, "Invalid Token", http.StatusUnauthorized);
				return;
			}


			claims := token.Claims.(jwt.MapClaims)
			ctx := context.WithValue(r.Context(), "email", claims["email"]);
			ctx = context.WithValue(ctx, "role", claims["role"]);
			next.ServeHTTP(w, r.WithContext(ctx));
		});
}


func RateLimitMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 2)
	
	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
		
			if (! limiter.Allow()) {

				log.Println("Rate Limiter exceeded the bucket size");
				response := map[string]interface{} {
					"message":"The API is at max-capacity. Please try later !",
					"success":false,
				}

				w.WriteHeader(http.StatusTooManyRequests);
				_ = json.NewEncoder(w).Encode(response);
			} else {
				next(w, r);
			}
		})
}