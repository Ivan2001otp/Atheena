package routers


import (
	"net/http"
	"github.com/gorilla/mux"

	"atheena/internals/middleware"
	"atheena/internals/handlers"
)


func RegisterAuthExitRouters(apiRouter *mux.Router) {

	apiRouter.Use(func(handler http.Handler) http.Handler {
		return middleware.RateLimitMiddleware(handler.ServeHTTP);
	})

	apiRouter.Use(func(handler http.Handler) http.Handler {
		return middleware.TokenMiddleware(handler);
	})

	apiRouter.HandleFunc("/logout",handlers.LogoutHandler).Methods("POST");
	apiRouter.HandleFunc("/delete_account", handlers.DeleteAccountHandler).Methods("POST");
	
}