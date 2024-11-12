package routes

import (
	"match_me_backend/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.GetLoginPage).Methods("GET")

	return router
}
