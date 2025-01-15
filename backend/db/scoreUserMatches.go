package db

import (
	"log"
	"match_me_backend/models"
	"match_me_backend/utils"
)

func CalculateUserDistance(matchID int, userID1, userID2 string) (float64, error) {
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

// Generates a match profile by combining all similar interests.
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

	// Extract distances
	for _, User1Interest := range user1Interests {
		for _, User2Interest := range user2Interests {
			if User1Interest == User2Interest {
				matchProfile = append(matchProfile, User1Interest)
			}
		}
	}

	matchScore := CalculateMatchProfile(matchProfile)

	// //TODO: This is in the wrong place the distance has not yet been set.
	// scoreModifier :=FilterByUserDistancePreference(userID1, userID2, user1Interests, user2Interests)
	// matchScore = matchScore * scoreModifier


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


//TODO if time permits - Refactor this function to use a switch and case statement and a helper function to return max distance
// Returns 1 if the match is valid, 0 if the match is invalid -1 if there is an error
func FilterByUserDistancePreference(distance float64, user1Interests, userInterests2 []models.Interests) int {

	var user1DistanceLimit int
	var user2DistanceLimit int

	// Gets the higher limit of the distance the user is willing to travel
	for _, interest := range user1Interests {
		// Checks to see if the user has marked the distance as preferred
		if interest.InterestName == UP_TO_ONE_HUNDRED {
			log.Println("User 1 distance limit is 100")
			user1DistanceLimit = 100
		}
		if interest.InterestName == ONE_HUNDRED_TO_FIVE_HUNDRED {
			log.Println("User 1 distance limit is 500")
			user1DistanceLimit = 500
		}
		if interest.InterestName == FIVE_HUNDRED_TO_ONE_THOUSAND {
			log.Println("User 1 distance limit is 1000")
			user1DistanceLimit = 1000
		}
		if interest.InterestName == ONE_THOUSAND_AND_BEYOND {
			log.Println("User 1 distance limit is 1001")
			user1DistanceLimit = 1001
		}
	}

	for _, interest := range userInterests2 {
		if interest.InterestName == UP_TO_ONE_HUNDRED {
			log.Println("User 2 distance limit is 100")
			user2DistanceLimit = 100
		}
		if interest.InterestName == ONE_HUNDRED_TO_FIVE_HUNDRED {
			log.Println("User 2 distance limit is 500")
			user2DistanceLimit = 500
		}
		if interest.InterestName == FIVE_HUNDRED_TO_ONE_THOUSAND {
			log.Println("User 2 distance limit is 1000")
			user2DistanceLimit = 1000
		}
		if interest.InterestName == ONE_THOUSAND_AND_BEYOND {
			log.Println("User 2 distance limit is 1001")
			user2DistanceLimit = 1001
		}
	}

	if user1DistanceLimit == 0 || user2DistanceLimit == 0 {
		log.Println("Error: Distance limits not properly set for one or both users")
		return -1
	}

	// If the distanceLimit is less than the distance between the users
	if float64(user1DistanceLimit) < distance || float64(user2DistanceLimit) < distance {
		return 0
	}
	// If the distanceLimit is more than the distance between the users, the match is valid
	return 1
}


func UpdateScoreByDistance(userID1, userID2 string, distance float64) error {
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
	// user1Interests is []models.Interests
	user2Interests := *user2InterestsPtr
	var matchScore int
	scoreModifier :=FilterByUserDistancePreference(distance, user1Interests, user2Interests)
	matchScore = matchScore * scoreModifier
	err = UpdateUserMatchScore(userID1, userID2, matchScore)
	return err
}