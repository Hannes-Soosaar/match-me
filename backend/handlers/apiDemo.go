package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"time"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetDemoUsers(w http.ResponseWriter, r *http.Request) {

	// Set environment variable for creating demo users true.
	var cycles int
	if db.InitDemoUsers() {
		userMatches, err := db.GetAllUserMatches()
		if err != nil {
			log.Println("Error getting user matches:", err)
		}
		for _, userMatch := range userMatches {
			db.CalculateUserDistance(userMatch.UserID1, userMatch.UserID2)
			db.CalculateMatchScore(userMatch.UserID1, userMatch.UserID2)
			if cycles == 100 {
				time.Sleep(1 * time.Second)
				log.Println("100 demo users spawned")
				cycles = 0
			}
			cycles++
		}
	}
	successMessage := "Demo bots spawned and are on the loose!"

	// set environment variable for creating demo users false.

	log.Println("Success!", successMessage)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successMessage)
}
