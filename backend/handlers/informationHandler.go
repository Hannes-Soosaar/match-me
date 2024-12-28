package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"strings"
	"time"
)

type PostUsernameRequest struct {
	Username string `json:"username"`
}

type PostCitynameRequest struct {
	City   string `json:"city"`
	Nation string `json:"country"`
	Region string `json:"state"`
}

type PostAboutRequest struct {
	About string `json:"about"`
}

type PostBirthdateRequest struct {
	Birthdate time.Time `json:"birthdate"`
}

func PostUsername(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	var body PostUsernameRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	if body.Username == "" {
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		log.Printf("Error: Username cannot be empty")
		return
	}

	err = db.SetUsername(currentUserID, body.Username)
	if err != nil {
		http.Error(w, "Error setting the username", http.StatusInternalServerError)
		log.Printf("Error setting the username: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Username successfully registered"})
}

func PostCity(w http.ResponseWriter, r *http.Request) {
	// Extract the Authorization header from the request
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}

	// Extract the token from the Authorization header
	token := strings.TrimPrefix(authHeader, "Bearer ")
	log.Println("JWT token extracted successfully.")

	// Extract the user ID from the JWT token
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	// Parse the request body to get city information
	var body PostCitynameRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	// Ensure that City, Longitude, and Latitude are provided and not empty
	if body.City == "" || body.Region == "" || body.Nation == "" {
		http.Error(w, "City details cannot be empty", http.StatusBadRequest)
		log.Printf("Error: City, Longitude, or Latitude cannot be empty")
		return
	}

	// Call the db function to set the city details for the user
	err = db.SetCity(currentUserID, body.Nation, body.Region, body.City)
	if err != nil {
		http.Error(w, "Error setting the City", http.StatusInternalServerError)
		log.Printf("Error setting the City: %v", err)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "City successfully registered"})
}

func PostAbout(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	var body PostAboutRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	if body.About == "" {
		http.Error(w, "About cannot be empty", http.StatusBadRequest)
		log.Printf("Error: About cannot be empty for userID=%s", currentUserID)
		return
	}

	err = db.SetAbout(currentUserID, body.About)
	if err != nil {
		http.Error(w, "Error setting the About", http.StatusInternalServerError)
		log.Printf("Error setting the About for userID=%s: %v", currentUserID, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "About successfully registered"})
}

func PostBirthdate(w http.ResponseWriter, r *http.Request) {
	// Extract Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}

	// Extract the token by trimming the 'Bearer ' prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")
	log.Println("JWT token extracted successfully.")

	// Extract the user ID from the token
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	// Parse the request body to get the birthdate
	var body PostBirthdateRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}

	// Validate that the Birthdate is not empty (optional, depending on your use case)
	if body.Birthdate.IsZero() {
		http.Error(w, "Birthdate cannot be empty", http.StatusBadRequest)
		log.Printf("Error: Birthdate cannot be empty for userID=%s", currentUserID)
		return
	}

	// Call the database function to update the birthdate for the user
	err = db.SetBirthdate(currentUserID, body.Birthdate)
	if err != nil {
		http.Error(w, "Error setting the birthdate", http.StatusInternalServerError)
		log.Printf("Error setting the birthdate for userID=%s: %v", currentUserID, err)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Birthdate successfully registered"})
}
