package routers

import (
	"net/http"
	"github.com/gorilla/mux"

	"atheena/internals/middleware"
	"atheena/internals/handlers"
)

func RegisterCommonRouters(apiRouter *mux.Router) {

	apiRouter.Use(func(next http.Handler) http.Handler {
		return middleware.RateLimitMiddleware(next.ServeHTTP)
	});

		
	apiRouter.HandleFunc("/login", handlers.LoginHandler).Methods("POST");
	apiRouter.HandleFunc("/register", handlers.RegisterHandler).Methods("POST");
	apiRouter.HandleFunc("/refresh-token", handlers.RefreshTokenHandler).Methods("POST");
}