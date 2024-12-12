package handlers

import (
	"encoding/json"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"os"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user RegLoginUser

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.Password == "" || user.Email == "" {
		sendErrorResponse(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	existingUser, err := db.GetUserByEmail(user.Email)
	time.Sleep(500 * time.Millisecond)
	if err != nil && err.Error() != "no user found with that email" {
		sendErrorResponse(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}

	if existingUser == nil || !auth.ComparePasswords(existingUser.Password, user.Password) {
		sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
		return
	}

	var jwtSecret = os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		sendErrorResponse(w, "JWT_SECRET not set in environment", http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(existingUser.ID)
	if err != nil {
		sendErrorResponse(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

//User-Match Handler
// Compare Scores
