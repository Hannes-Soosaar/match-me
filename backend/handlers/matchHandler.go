package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"net/http"
)

// Handle match requests from the front end

type MatchRequest struct {
	MatchId int `json:"match_id"`
}

func RemoveMatch(w http.ResponseWriter, r *http.Request) {
	var payload MatchRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	matchId := payload.MatchId

	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}
	// checks if the user is logged in
	_, err = GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	successMessage, err := db.UpdateUserMatchStatus(matchId, db.REMOVED)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func RequestMatch(w http.ResponseWriter, r *http.Request) {

	var payload MatchRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	matchId := payload.MatchId

	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}

	_, err = GetCurrentUserID(r)

	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	successMessage, err := db.UpdateUserMatchStatus(matchId, db.REQUESTED)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func ConfirmMatch(w http.ResponseWriter, r *http.Request) {

	var payload MatchRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	matchId := payload.MatchId

	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}

	_, err = GetCurrentUserID(r)

	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	successMessage, err := db.UpdateUserMatchStatus(matchId, db.CONNECTED)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func BlockMatch(w http.ResponseWriter, r *http.Request) {

	var payload MatchRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	matchId := payload.MatchId

	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}

	_, err = GetCurrentUserID(r)

	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	successMessage, err := db.UpdateUserMatchStatus(matchId, db.BLOCKED)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

// is online ?
type MatchResponse struct {
	MatchID                int    `json:"match_id"`
	MatchScore             int    `json:"match_score"`
	Status                 string `json:"status"`
	MatchedUserName        string `json:"matched_user_name"`
	MatchedUserPicture     string `json:"matched_user_picture"`
	MatchedUserDescription string `json:"matched_user_description"`
	MatchedUserLocation    string `json:"matched_user_location"`
}

func GetMatches(w http.ResponseWriter, r *http.Request) {

	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	// Based on what the input is here we can filter the output to the front end.
	userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	log.Println(`userMatches`, userMatches)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}

	var matches []MatchResponse
	var match MatchResponse
	for _, userMatch := range userMatches {
		matchProfile, err := db.GetUserInformation(userMatch.UserID2)
		if err != nil {
			log.Println("Error getting user information:", err)
		}
		match.MatchID = userMatch.ID
		match.Status = userMatch.Status
		match.MatchScore = userMatch.MatchScore
		match.MatchedUserName = matchProfile.Username
		match.MatchedUserPicture = matchProfile.Picture
		match.MatchedUserDescription = matchProfile.About
		match.MatchedUserLocation = matchProfile.Nation
		matches = append(matches, match)
		log.Println("Match in loop:", match)
	}

	log.Println("Matches:", matches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matches)
}

func GetConnections(w http.ResponseWriter, r *http.Request) {

	userID1, err := GetCurrentUserID(r)

	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	connectionsUUIDs, err := db.GetConnectionsID(userID1)
	if err != nil {
		log.Println("Error getting connections UUID from db:", err)
	}

	connectionsIDs, err := db.GetUserIDfromUUIDarray(connectionsUUIDs)
	if err != nil {
		log.Println("Error getting connections ID from db:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(connectionsIDs)
}
