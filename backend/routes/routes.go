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
	router.HandleFunc("/users/{id}/profile", handlers.GetUserProfileHandler).Methods("GET")
	//router.HandleFunc("/users/{id}/bio", handlers.GetUserBioHandler).Methods("GET") // TEST
	router.HandleFunc("/me", handlers.GetLightCurrentUserHandler).Methods("GET")
	router.HandleFunc("/me/profile", handlers.GetCurrentUserHandler).Methods("GET")
	//router.HandleFunc("/me/bio", handlers.GetMeBioHandler).Methods("GET") // TEST

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
	router.HandleFunc("/recommendations", handlers.GetReccomendationsHandler).Methods("GET") // Test

	/*
		EXTERNAL API ROUTES

		Create documentation for the following routes:

		DONE:   /users/{id}: which returns the user's name and link to the profile picture.
		DONE:   /users/{id}/profile: which returns the users "about me" type information.

		None of them must return authentication-related data.

		/users/{id}/bio: which returns the users biographical data (the data used to power recommendations).

		  DONE AND IN USER:  /me: which is a shortcut to /users/{id}  for the authenticated user. You should also implement

		TODO: router.HandleFunc("/me/bio", handlers.GetCurrentUserHandler).Methods("GET")
		TODO: router.HandleFunc("/me/profile", handlers.GetCurrentUserHandler).Methods("GET")

		TODO:  /recommendations: which returns a maximum of 10 recommendations, containing only the id and nothing else.

		TODO: /connections: which returns a list connected profiles, containing only the id and nothing else.

		TODO: All of the responses for /users data must also contain the id.

		TODO: If the id is not found, or the user does not have permission to view that profile, it must return HTTP404.
	*/

	// chat routes
	router.HandleFunc("/ws", handlers.WebsocketHandler)

	return router
}
