package routes

import (
	"match_me_backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	fileDirectory := "../frontend/src/components/Assets" // Modify with your actual directory

	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(fileDirectory))))

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/me", handlers.GetCurrentUserHandler).Methods("GET")
	router.HandleFunc("/test", handlers.GetTestResultHandler).Methods("GET")
	router.HandleFunc("/userInterests", handlers.GetUserInterests).Methods("GET")
	router.HandleFunc("/userInterest", handlers.UpdateUserInterest).Methods("POST")
	router.HandleFunc("/interests", handlers.GetInterests).Methods("GET")
	router.HandleFunc("/username", handlers.PostUsername).Methods("POST")
	// router.HandleFunc("/match", handlers.UserMatches).Methods("GET")
	router.HandleFunc("/city", handlers.PostCity).Methods("POST")
	router.HandleFunc("/about", handlers.PostAbout).Methods("POST")
	router.HandleFunc("/birthdate", handlers.PostBirthdate).Methods("POST")
	router.HandleFunc("/picture", handlers.PostProfilePictureHandler).Methods("POST")

	router.HandleFunc("/browserlocation", handlers.BrowserHandler).Methods("POST")

	return router
}
