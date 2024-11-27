package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"match_me_backend/db"
	"net/http"

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

<<<<<<< HEAD
func GetMeHandler(w http.ResponseWriter, r *http.Request) {

}
=======
func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request){
	queryToken := r.URL.Query().Get("token")
	if queryToken != "" {
		fmt.Println("Token from query:", queryToken)
	}
	fmt.Println("Running the function")
}


>>>>>>> refs/remotes/origin/main
