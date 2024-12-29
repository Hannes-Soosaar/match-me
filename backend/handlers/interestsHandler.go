package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"net/http"
)

func GetUserInterests(w http.ResponseWriter, r *http.Request) {

	userID, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error:", err)
	}	
	// need to get the user ID.
	 userInterestIDs,err := db.GetAllUserInterestIDs(userID);

	 if err != nil {	
		log.Println("Error in GetUserInterestsId's", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("GetUserInterests")
}

func GetInterests(w http.ResponseWriter, r *http.Request) {
	// allCategories, err := db.GetAllCategories()
	interests,err := db.GetAllInterest()

	if err != nil {
		log.Println("Error getting interest response ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("GetInterests")
}