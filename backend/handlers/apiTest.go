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

	userID := "d5d084c8-927a-4c55-81b7-fe00496e1a68" // user id for a@a.com
	// userID2 := "0ee5d527-351b-4be5-ade4-7e93614a259c" // user id for hsoosaar@gmail.com

	// err := db.AddUserMatch(userID, userID2)
	// score , err := utils.CalculateMatchScore(userID, userID2)
	// if err != nil {
	// 	log.Println("Error calculating match score", err)
	// }

	// log.Println("Match score: ", score)

	userMatches, err := db.GetAllUserMatchesByUserId(userID)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}

	log.Println("User matches are: ", userMatches)

	// log.Println("users interest are: ", usersInterest)

	// usersInterest2, err := db.GetAllUserInterest(userID2)
	// if err != nil {
	// 	log.Printf(" the error is , %s", err)
	// }
	// log.Println("users interest are: ", usersInterest2)

	// 	// allCategories, err := db.GetAllCategories()
	// 	// allInterests, err := db.GetAllInterest()
	// 	interestResponseBody, err := db.GetInterestResponseBody()

	// 	if err != nil {
	// 		log.Printf(" the error is , %s", err)
	// 	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Println("Test results logged")

	json.NewEncoder(w).Encode(userMatches)
}
