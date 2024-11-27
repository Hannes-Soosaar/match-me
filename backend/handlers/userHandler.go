package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"match_me_backend/db"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UserResponse struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user UserResponse

	err := db.DB.QueryRow(`	SELECT u.id, p.username, p.profile_picture
							FROM users u
							JOIN profiles p ON u.id = p.user_id
							WHERE u.id = $1`, userID).Scan(&user.ID, &user.Username, &user.ProfilePicture)

	if err == sql.ErrNoRows {
		http.Error(w, "User/Profile not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Printf("error: %v\n", err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func GetMeHandler(w http.ResponseWriter, r *http.Request) {

}

func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UserResponse
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	fmt.Println("Extracted JWT:", token)
		user.ID = 9001;
		user.Username ="test"
		user.ProfilePicture ="my picture"
	w.Write([]byte("Access granted."))
	json.NewEncoder(w).Encode(user)
}
