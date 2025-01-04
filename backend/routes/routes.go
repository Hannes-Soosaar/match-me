package routes

import (
	"match_me_backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	fileDirectory := "../frontend/src/components/Assets/ProfilePictures"

	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(fileDirectory))))
	// user routes
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/me", handlers.GetCurrentUserHandler).Methods("GET")
	router.HandleFunc("/test", handlers.GetTestResultHandler).Methods("GET")

	// Profile routes
	router.HandleFunc("/userInterests", handlers.GetUserInterests).Methods("GET")
	router.HandleFunc("/userInterest", handlers.UpdateUserInterest).Methods("POST")
	router.HandleFunc("/interests", handlers.GetInterests).Methods("GET")
	router.HandleFunc("/username", handlers.PostUsername).Methods("POST")
	router.HandleFunc("/city", handlers.PostCity).Methods("POST")
	router.HandleFunc("/about", handlers.PostAbout).Methods("POST")
	router.HandleFunc("/birthdate", handlers.PostBirthdate).Methods("POST")
	router.HandleFunc("/picture", handlers.PostProfilePictureHandler).Methods("POST")

	router.HandleFunc("/browserlocation", handlers.BrowserHandler).Methods("POST")

	// match routes
	router.HandleFunc("/matches", handlers.GetMatches).Methods("GET") // get 15 all matches
	router.HandleFunc("/matches/request", handlers.RequestMatch).Methods("PUT")
	router.HandleFunc("/matches/connect", handlers.ConfirmMatch).Methods("PUT")
	router.HandleFunc("/matches/block", handlers.BlockMatch).Methods("PUT")
	router.HandleFunc("/matches/remove", handlers.RemoveMatch).Methods("PUT")
	router.HandleFunc("/connections", handlers.GetConnections).Methods("GET")

	// chat routes
	router.HandleFunc("/ws", handlers.WebsocketHandler)

	return router
}
