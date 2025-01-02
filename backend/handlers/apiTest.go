package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	// "match_me_backend/utils"
	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {
	// utils.InitDemoUsers()
	// userID := "d5d084c8-927a-4c55-81b7-fe00496e1a68" // user id for a@a.com
	userID1:= "cea71d69-d41f-4b76-a2af-c1fb3bbc35b6" // user id for a@a.com
	db.AddUserMatchForAllExistingUsers(userID1)




	userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}

	log.Println("User matches are: ", userMatches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("Test results logged")

	json.NewEncoder(w).Encode(userMatches)
}
