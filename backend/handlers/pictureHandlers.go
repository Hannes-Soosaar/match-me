package handlers

import (
	"fmt"
	"io"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PostProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	currentUserID, err := auth.ExtractUserIDFromToken(token)

	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}

	file, _, err := r.FormFile("profilePic") // "profilePic" should be the name of the form field
	if err != nil {
		http.Error(w, "Unable to extract file", http.StatusBadRequest)
		log.Printf("Error extracting file: %v", err)
		return
	}
	defer file.Close()

	rand.Seed(time.Now().UnixNano())
	randomFileName := fmt.Sprintf("%s_%d.jpeg", currentUserID, rand.Intn(1000000))

	dir := "../frontend/src/components/Assets" // Change this to the desired directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			http.Error(w, "Unable to create directory", http.StatusInternalServerError)
			log.Printf("Error creating directory: %v", err)
			return
		}
	}

	filePath := filepath.Join(dir, randomFileName)
	destFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		log.Printf("Error creating file: %v", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, "Unable to save file content", http.StatusInternalServerError)
		log.Printf("Error copying file content: %v", err)
		return
	}

	err = db.SetPicturePath(currentUserID, filePath)
	if err != nil {
		http.Error(w, "Error setting the username", http.StatusInternalServerError)
		log.Printf("Error setting the username: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Profile picture uploaded successfully! Saved as %s", randomFileName)))
}
