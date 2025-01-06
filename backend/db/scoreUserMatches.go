package db

import (
	"log"
	"match_me_backend/models"
	"match_me_backend/utils"
)

func CalculateUserDistance(matchID int,userID1, userID2 string) (float64, error) {
	user1, err := GetUserByID(userID1)
	if err != nil {
		log.Println("Error getting user 1", err)
	}
	user2, err := GetUserByID(userID2)
	if err != nil {
		log.Println("Error getting user 2", err)
	}
	distance := utils.GetDistanceBetweenTwoPointsOnEarth(user1.Latitude, user1.Longitude, user2.Latitude, user2.Longitude)
	err = UpdateMatchDistance(matchID, distance)
	if err != nil {
		log.Println("Error updating match distance", err)
	}
	return distance, err
}	

// Generates the match score for a new user The interest must already exist!
func CalculateMatchScore(userID1, userID2 string) (int, error) {

	user1InterestsPtr, err := GetAllUserInterest(userID1)

	if err != nil {
		log.Println("Error getting user 1 interest", err)
	}
	// user1Interests is []models.Interests
	user1Interests := *user1InterestsPtr
	user2InterestsPtr, err := GetAllUserInterest(userID2)

	if err != nil {
		log.Println("Error getting user 2 interest", err)
	}
	user2Interests := *user2InterestsPtr

	var matchProfile []models.Interests

	for _, User1Interest := range user1Interests {
		for _, User2Interest := range user2Interests {
			if User1Interest == User2Interest {
				matchProfile = append(matchProfile, User1Interest)
			}
		}
	}

	matchScore := CalculateMatchProfile(matchProfile)
	err = UpdateUserMatchScore(userID1, userID2, matchScore)
	return matchScore, err
}

func CalculateMatchProfile(matchProfile []models.Interests) int {

	score := 0
	genreCount := 0
	playStyleCount := 0
	platformCount := 0
	communicationCount := 0
	goalsCount := 0
	sessionCount := 0
	vibeCount := 0
	languageCount := 0

	// 5 is a good score meaning there is a least one match per category
	// excluding language, platform and communication

	for _, interest := range matchProfile {
		if interest.CategoryID == GENRE {
			genreCount += 1
			score += 1
		}
		if interest.CategoryID == PLAY_STYLE {
			playStyleCount += 1
			score += 1
		}
		if interest.CategoryID == PLATFORM {
			platformCount += 1
		}
		if interest.CategoryID == COMMUNICATION {
			communicationCount += 1
		}
		if interest.CategoryID == GOALS {
			goalsCount += 1
			score += 1
		}
		if interest.CategoryID == SESSION {
			sessionCount += 1
			score += 1
		}
		if interest.CategoryID == VIBE {
			vibeCount += 1
			score += 1
		}
		if interest.CategoryID == LANGUAGE {
			languageCount += 1
		}
	}

	/* THE BELOW SECTION HAS BEEN REMOVED BECAUSE OF ALMOST NO MATCHES */
	// If the user has less than 1 interest in any of these category, the match score is set to 0
	if languageCount < 1 || platformCount < 1 || communicationCount < 1 {
		return 0
	}

	// the score is only derived from the number of interests in the categories of Genre, PlayStyle, Goals, Session and Vibe

	return score

}
