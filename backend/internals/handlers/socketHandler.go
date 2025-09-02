package handlers

import (
	_envConfig "atheena/internals/config"
	_sockets "atheena/internals/database/websockets"
	"context"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true;
	},
}


func generateUniqueID() string {
	return uuid.New().String()
}

func getUserIdFromContext(ctx context.Context) primitive.ObjectID {
	userIdStr, ok := ctx.Value("user_id").(string)

	if !ok {
		log.Println("❌Failed to get user_id from context. Seems to be its not string datatype");
		return primitive.NilObjectID
	}

	userID, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		log.Printf("❌Failed to convert user_id to ObjectID: %v", err);
		return primitive.NilObjectID;
	}

	return userID
}

func ServeWs(hub *_sockets.Hub, w http.ResponseWriter, r *http.Request) {
	
	 // Get token from URL parameters
    token := r.URL.Query().Get("token")
    
    // Create a new context with the token
    ctx := r.Context()
    if token != "" {
        // Parse the token
        parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
            secretKey := _envConfig.LoadEnvConfig().JWT_Secret
            return []byte(secretKey), nil
        })

        if err != nil {
            log.Printf("❌ Invalid token: %v", err)
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
            // Add claims to context
			log.Println("The claims are ", claims);
            ctx = context.WithValue(ctx, "email", claims["email"])
            ctx = context.WithValue(ctx, "role", claims["role"])
            ctx = context.WithValue(ctx, "user_id", claims["user_id"])
        } else {
            log.Println("❌ Invalid token claims")
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }
    }
	
	
	
	conn, err := upgrader.Upgrade(w, r, nil);
	if err != nil {
		log.Printf("❌ Failed to upgrade connection: %v", err)
		log.Println(err)
		return;
	}
	
	userID := getUserIdFromContext(ctx)
	if userID == primitive.NilObjectID {
        log.Println("❌ No valid userID found in context")
        conn.Close()
        return
    }
	log.Printf("✅ User connected with ID: %s", userID.Hex())


	client := & _sockets.Client{
		ID: generateUniqueID(),
		Hub: hub,
		Conn: conn,
		UserID: userID,
		Send : make(chan []byte, 256),
	}
	
	log.Printf("📝 Registering client with ID: %s", client.ID)
    
	client.Hub.Register <- client
	go client.WritePump()
	go client.ReadPump()
}