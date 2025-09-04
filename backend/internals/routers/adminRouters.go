package routers


import (
	"net/http"
	"github.com/gorilla/mux"

	"atheena/internals/middleware"
	"atheena/internals/handlers"
)


func RegisterAdminRouters(apiRouter *mux.Router) {
	apiRouter.Use(func(handler http.Handler) http.Handler{
		return middleware.RateLimitMiddleware(handler.ServeHTTP)
	});


	apiRouter.Use(func(handler http.Handler) http.Handler {
		return middleware.TokenMiddleware(handler);
	})

	apiRouter.HandleFunc("/upsert_supervisor",handlers.AddOrUpdateSupervisor).Methods("POST");
	apiRouter.HandleFunc("/delete_supervisor", handlers.DeleteSupervisor).Methods("POST");
	apiRouter.HandleFunc("/get_supervisors", handlers.FetchAllSupervisor).Methods("GET");

	apiRouter.HandleFunc("/add_construction_site", handlers.AddConstructionSite).Methods("POST");
	apiRouter.HandleFunc("/add_warehouse", handlers.AddNewWarehouse).Methods("POST");

	apiRouter.HandleFunc("/get_warehouses/{admin_id}", handlers.GetAllWarehouseByAdminId).Methods("GET");
	apiRouter.HandleFunc("/get_construction_sites/{admin_id}", handlers.FetchConstructionSitebyAdminId).Methods("GET");

	apiRouter.HandleFunc("/add_inventory", handlers.AddInventoryItem).Methods("POST");
	apiRouter.HandleFunc("/get_inventory", handlers.FetchInventoryByWarehouse).Methods("GET");

	apiRouter.HandleFunc("/get_all_logs", handlers.FetchAllLogs).Methods("GET");

	// order tracking system endpoints
	apiRouter.HandleFunc("/approve_order", handlers.InteractOrderApproval).Methods("GET");
}