package routers

import (
	"net/http"
	"github.com/gorilla/mux"

	"atheena/internals/middleware"
	"atheena/internals/handlers"
)

func RegisterSupervisorRouters(apiRouter *mux.Router) {
	apiRouter.Use(func(handler http.Handler) http.Handler{
		return middleware.RateLimitMiddleware(handler.ServeHTTP)
	});

	
	apiRouter.HandleFunc("/send_approval", handlers.AskForApproval).Methods("POST");
}