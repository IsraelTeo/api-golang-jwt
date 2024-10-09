package route

import (
	"github.com/IsralTeo/api-go-jwt/controller"
	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	routes := mux.NewRouter()
	api := routes.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", controller.InitRoute).Methods("GET")
}
