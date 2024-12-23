package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UserResponse struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}

// TODO: move to  db user_queries
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := db.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User/Profile not found", http.StatusNotFound)
			log.Printf("User/Profile not found for uuid=%s: %v", userID, err)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error fetching user/profile for uuid=%s: %v", userID, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {

}

func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *models.ProfileInformation

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	log.Println("JWT token extracted successfully.")

	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	user, err = db.GetUserInformation(currentUserID)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			log.Printf("User with ID %v not found", currentUserID)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error fetching user information: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
