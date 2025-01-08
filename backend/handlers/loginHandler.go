package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user RegLoginUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Invalid login request: %v", err)
		return
	}

	if user.Password == "" || (user.Email == "" && user.Username == "") {
		sendErrorResponse(w, "Password and either Email or Username are required", http.StatusBadRequest)
		log.Println("Password and either Email or Username are required")
		return
	}

	var existingUser *models.User
	if user.Username != "" {
		existingUser, err = db.GetUserByUsername(user.Username)
	} else {
		existingUser, err = db.GetUserByEmail(user.Email)
	}

	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
			log.Println("Email or password is incorrect")
			return
		}
		sendErrorResponse(w, "Error checking user existence", http.StatusInternalServerError)
		log.Printf("Error checking user existence: %v", err)
		return
	}

	if existingUser == nil || !auth.ComparePasswords(existingUser.PasswordHash, user.Password) {
		sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
		log.Println("Email or password is incorrect")
		return
	}

	token, err := auth.GenerateJWT(existingUser.ID)
	if err != nil {
		sendErrorResponse(w, "Error generating JWT", http.StatusInternalServerError)
		log.Printf("Error generating JWT: %v", err)
		return
	}

	//is it redundant  to check if token is empty?
	if token == "" {
		sendErrorResponse(w, "No token", http.StatusInternalServerError)
		log.Println("Error token is empty")
		return
	}

	
	err = db.UpdateMatchScoreForUser(existingUser.ID)
	
	if err != nil {
		log.Println("Error updating all user scores", err)
	}

	err = db.SetUserOnlineStatus(existingUser.ID, true)

	if err != nil {
		log.Println("Error setting user online status", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := GetCurrentUserID(r)
	if err != nil {
		sendErrorResponse(w, "Error getting user Id from token", http.StatusUnauthorized)
		log.Println("Error getting user Id from token:", err)
		return
	}
	err = db.SetUserOnlineStatus(userId, false)
	if err != nil {
		sendErrorResponse(w, "Error setting user offline", http.StatusInternalServerError)
		log.Println("Error setting user offline:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "safely logged out!"})
}

