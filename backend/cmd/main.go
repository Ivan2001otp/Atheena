package main

import (
	envConfig "atheena/internals/config"
	_mongo "atheena/internals/database/mongoV2"
	_websockets "atheena/internals/database/websockets"
	_websocketH "atheena/internals/handlers"
	
	"atheena/internals/handlers"
	"atheena/internals/routers"

	_util "atheena/internals/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)


func setUpConfig() {
	env := envConfig.LoadEnvConfig()

   if env != nil {
		log.Println("✅ Environment Variables loaded successfully");
   } else {
	log.Fatal("❌ Could not load Environment variables.")
   }


   _ , err := _mongo.GetMongoClient()
   if err != nil {
	log.Fatal("❌ Could not set up MongoDB !");
   } else {
	log.Println("MongoDB is set up successfully !");
   }

   _util.Init()
}

func main() {
	hub:= _websockets.GetSocketHub()
	
	setUpConfig()
	// define the cors
	corsOptions := cors.New(cors.Options{
		AllowedMethods : []string{"GET", "DELETE", "POST", "OPTIONS", "PUT"},
		AllowCredentials : true,
		AllowedHeaders : []string{"Authorization", "Content-Type"},
		AllowedOrigins: []string{"http://localhost:5173"},
	})

	
	// establish routers
	mainRouter := mux.NewRouter()


	// websocket
	mainRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		_websocketH.ServeWs(hub, w, r)
	})

	// Custom Handle 404 error .
	mainRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound);

			json.NewEncoder(w).Encode(map[string]interface{}{
				"success":false,
				"message":"The API endpoint you are trying to reach does not exist",
			})
	})

	//auth
	mainRouter.HandleFunc("/api/v1/login", handlers.LoginHandler).Methods("POST");
	mainRouter.HandleFunc("/api/v1/register", handlers.RegisterHandler).Methods("POST");
	mainRouter.HandleFunc("/api/v1/refresh-token", handlers.RefreshTokenHandler).Methods("POST");

	commonRouters := mainRouter.PathPrefix("/api/v1").Subrouter()
	exitRouters := mainRouter.PathPrefix("/api/v1").Subrouter();

	adminRouters := mainRouter.PathPrefix("/api/v1").Subrouter();
	supervisorRouters := mainRouter.PathPrefix("/api/v1").Subrouter();

	routers.RegisterCommonRouters(commonRouters);
	routers.RegisterAuthExitRouters(exitRouters);
	routers.RegisterAdminRouters(adminRouters);
	routers.RegisterSupervisorRouters(supervisorRouters);

	// setting cors options
	handler := corsOptions.Handler(mainRouter)
	port := envConfig.LoadEnvConfig().Port
	if port == "" {
		port = "8080";
	}

	

	log.Println("Backend listening at port : "+port);
	log.Fatal(http.ListenAndServe(":"+port, handler));
}