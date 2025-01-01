package utils

import (
	"log"
	"match_me_backend/db"
	"match_me_backend/models"
)

// Read in all the matches
// Find matches with no scores
// Score

// Generates the match score for a new user
func CalculateMatchScore(userID1, userID2 string) (int,error) {

	user1InterestsPtr, err := db.GetAllUserInterest(userID1)
	if err != nil {
		log.Println("Error getting user 1 interest", err)
	}
	// user1Interests is []models.Interests
	user1Interests := *user1InterestsPtr
	user2InterestsPtr, err := db.GetAllUserInterest(userID2)

	if err != nil {
		log.Println("Error getting user 2 interest", err)
	}
	user2Interests := *user2InterestsPtr
	log.Println("User 1 interests: ", user1Interests)
	log.Println("User 2 interests: ", user2Interests)

	var matchProfile []models.Interests

	for _, User1Interest := range user1Interests {

		for _, User2Interest := range user2Interests {

			// handel exceptions where the score must be zero
			if User1Interest == User2Interest {
				log.Println("Match found")
				log.Println("User 1 interest: ", User1Interest)
				log.Println("User 2 interest: ", User2Interest)
				//Create a match profile for the two users.
				matchProfile = append(matchProfile, User1Interest)
			}
		}
	}
	matchScore := CalculateMatchProfile(matchProfile)

//TODO this is temporary until we have the full update path determined
	err = db.UpdateUserMatchScore(userID1, userID2, matchScore)
	return matchScore, err
}



// If the location criteria is not met, the match score is set to 0
func ZeroMatchScore(currentUserID, userID2 string) {
	err := db.UpdateUserMatchScore(currentUserID, userID2, 0)
	if err != nil {
		log.Println("Error updating match score: ", err)
	}
}

// Update the match score in the database use when a user updates their profile
func UpdatedMatchScore() {

}

func CalculateMatchProfile(matchProfile []models.Interests) (int) {
	score := 0
	genreCount := 0
	playStyleCount := 0
	platformCount := 0
	communicationCount := 0
	goalsCount := 0
	sessionCount := 0
	vibeCount := 0
	languageCount := 0

	for _, interest := range matchProfile {
		if interest.CategoryID == db.GENRE {
			genreCount += 1
			score += 1
		}
		if interest.CategoryID == db.PLAY_STYLE {
			playStyleCount += 1
			score += 1
		}
		if interest.CategoryID == db.PLATFORM {
			platformCount += 1
			score += 1
		}
		if interest.CategoryID == db.COMMUNICATION {
			communicationCount += 1
			score += 1
		}
		if interest.CategoryID == db.GOALS {
			goalsCount += 1
			score += 1
		}
		if interest.CategoryID == db.SESSION {
			sessionCount += 1
			score += 1
		}
		if interest.CategoryID == db.VIBE {
			vibeCount += 1
			score += 1
		}
		if interest.CategoryID == db.LANGUAGE {
			languageCount += 1
			score += 1
		}
	}

	// If the user has less than 1 interest in any of these category, the match score is set to 0
	if languageCount < 1 || platformCount < 1 || communicationCount < 1 {
		return 0
	}
	// the score is only derived from the number of interests in the categories of Genre, PlayStyle, Goals, Session and Vibe
	score = score -languageCount- platformCount - communicationCount 
	return score

}
