package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetDemoUsers(w http.ResponseWriter, r *http.Request) {


	if db.InitDemoUsers() {
		userMatches, err := db.GetAllUserMatches()
		if err != nil {
			log.Println("Error getting user matches:", err)
		}

		for _, userMatch := range userMatches {
			db.CalculateMatchScore(userMatch.UserID1, userMatch.UserID2)
			db.CalculateUserDistance(userMatch.ID , userMatch.UserID1, userMatch.UserID2)
		}

	}
	successMessage := "Demo bots spawned and are on the loose!"

	log.Println("Success!", successMessage)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successMessage)
}
