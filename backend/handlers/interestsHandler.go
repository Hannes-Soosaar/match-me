package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"net/http"
)

func GetUserInterests(w http.ResponseWriter, r *http.Request) {

	userID, err := GetCurrentUserID(r)

	if err != nil {
		log.Println("Error:", err)
	}
	
	userInterestIDs, err := db.GetAllUserInterestIDs(userID)

	if err != nil {
		log.Println("Error in GetUserInterestsId's", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userInterestIDs)
}

func GetInterests(w http.ResponseWriter, r *http.Request) {

	interests, err := db.GetInterestResponseBody()

	if err != nil {
		log.Println("Error getting interest response ", err)
	}
	log.Println("Interests:", interests)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(interests)
}
