package routes

import (
	"match_me_backend/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/me", handlers.GetCurrentUserHandler).Methods("GET")

	return router
}
