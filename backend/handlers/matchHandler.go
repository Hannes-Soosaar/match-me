package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
	"sort"
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

// returns  10 new matches for the user to see ordered by the best match score
func GetMatches(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	// userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	userMatches, err := db.GetTenNewMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	var matches []MatchResponse
	var match MatchResponse
	var matchProfile *models.ProfileInformation
	for _, userMatch := range userMatches {
		// Displays correctly the matched Profile user data
		if userMatch.UserID2 == userID1 {
			matchProfile, err = db.GetUserInformation(userMatch.UserID1)
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			matchProfile, err = db.GetUserInformation(userMatch.UserID2)
		}

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
	}

	log.Println("Matches:", matches)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matches)
}

func GetRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get all user matches for the current user
	userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(`userMatches`, userMatches)
	// Transform userMatches into MatchResponse objects
	var matches []MatchResponse
	for _, userMatch := range userMatches {
		matchProfile, err := db.GetUserInformation(userMatch.UserID2)
		if err != nil {
			log.Println("Error getting user information:", err)
			continue // Skip this match if there's an error
		}
		match := MatchResponse{
			MatchID:                userMatch.ID,
			Status:                 userMatch.Status,
			MatchScore:             userMatch.MatchScore,
			MatchedUserName:        matchProfile.Username,
			MatchedUserPicture:     matchProfile.Picture,
			MatchedUserDescription: matchProfile.About,
			MatchedUserLocation:    matchProfile.Nation,
		}
		matches = append(matches, match)
	}
	// Sort matches by MatchScore in descending order
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].MatchScore > matches[j].MatchScore
	})
	// Take the top 10 matches
	if len(matches) > 10 {
		matches = matches[:10]
	}
	// Encode the matches as JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(connectionsUUIDs)
}

type BuddiesResponse struct {
	MatchID                int    `json:"match_id"`
	MatchScore             int    `json:"match_score"`
	Status                 string `json:"status"`
	MatchedUserName        string `json:"matched_user_name"`
	MatchedUserPicture     string `json:"matched_user_picture"`
	MatchedUserDescription string `json:"matched_user_description"`
	MatchedUserLocation    string `json:"matched_user_location"`
}

// Returns the user's buddies who are connected
func GetBuddies(w http.ResponseWriter, r *http.Request) {

	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	userBuddies, err := db.GetAllConnectedMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user connected matches:", err)
	}

	log.Println(`userBuddies`, userBuddies);

	var buddies []BuddiesResponse
	var buddy BuddiesResponse
	var buddyProfile *models.ProfileInformation
	for _, userMatch := range userBuddies {
		// Displays correctly the matched Profile user data
		if userMatch.UserID2 == userID1 {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID1)
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID2)
		}

		if err != nil {
			log.Println("Error getting user information:", err)
		}
		buddy.MatchID = userMatch.ID
		buddy.Status = userMatch.Status
		buddy.MatchScore = userMatch.MatchScore
		buddy.MatchedUserName = buddyProfile.Username
		buddy.MatchedUserPicture = buddyProfile.Picture
		buddy.MatchedUserDescription = buddyProfile.About
		buddy.MatchedUserLocation = buddyProfile.Nation
		buddies = append(buddies, buddy)
	}

	log.Println("Buddies:", buddies)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buddies)
}



type BuddyProfile struct {
	MatchID                int    `json:"match_id"`
	MatchScore             int    `json:"match_score"`
	Status                 string `json:"status"`
	IsOnline           	   bool   `json:"is_online"`
	MatchedUserName        string `json:"matched_user_name"`
	MatchedUserPicture     string `json:"matched_user_picture"`
	MatchedUserDescription string `json:"matched_user_description"`
	MatchedUserLocation    string `json:"matched_user_location"`
}
// Will get the match ID  and return the buddy profile.
func GetBuddyProfile(w http.ResponseWriter, r *http.Request) {

	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}

	
	userBuddies, err := db.GetAllConnectedMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user connected matches:", err)
	}

	log.Println(`userBuddies`, userBuddies);

	var buddies []BuddiesResponse
	var buddy BuddiesResponse
	var buddyProfile *models.ProfileInformation
	for _, userMatch := range userBuddies {
		// Displays correctly the matched Profile user data
		if userMatch.UserID2 == userID1 {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID1)
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID2)
		}

		if err != nil {
			log.Println("Error getting user information:", err)
		}
		buddy.MatchID = userMatch.ID
		buddy.Status = userMatch.Status
		buddy.MatchScore = userMatch.MatchScore
		buddy.MatchedUserName = buddyProfile.Username
		buddy.MatchedUserPicture = buddyProfile.Picture
		buddy.MatchedUserDescription = buddyProfile.About
		buddy.MatchedUserLocation = buddyProfile.Nation
		buddies = append(buddies, buddy)
	}

	log.Println("Buddies:", buddies)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buddies)
}
