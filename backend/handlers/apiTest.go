package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {

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
	userID := "RAN this"
	// user id for a@a.com

	// db.AddUserMatchForAllExistingUsers(userID1)

	// userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	// if err != nil {
	// 	log.Println("Error getting user matches:", err)
	// }

	log.Println("User matches are: ", userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("Test results logged")

	json.NewEncoder(w).Encode(userID)
}
