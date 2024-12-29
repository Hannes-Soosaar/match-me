package db

import (
	"log"
	"match_me_backend/models"
)

type UserInterestResponse struct {
	CategoryName string             `json:"category"`
	CategoryID   int                `json:"category_id"`
	Interest     []models.Interests `json:"interests"`
}

func GetInterestResponseBody() (*[]UserInterestResponse, error) {
	categories, err := GetAllCategories()

	if err != nil {
		log.Println("Error getting all categories")
		return nil, err
	}
	interests, err := GetAllInterest()

	if err != nil {
		log.Println("Error getting all interests")
		return nil, err
	}

	interestMap := make(map[int][]models.Interests)
	for _, interest := range *interests {
		interestMap[interest.CategoryID] = append(interestMap[interest.CategoryID], interest)
	}

	var userInterestResponses []UserInterestResponse
	for _, category := range *categories {
		response := UserInterestResponse{
			CategoryName: category.CategoryName,
			CategoryID:   category.ID,
			Interest:     interestMap[category.ID],
		}
		userInterestResponses = append(userInterestResponses, response)
	}
	return &userInterestResponses, nil
}

// Gets all the interest
func GetAllInterest() (*[]models.Interests, error) {
	query := "SELECT * FROM interests"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interests []models.Interests
	for rows.Next() {
		var interest models.Interests
		err = rows.Scan(&interest.ID, &interest.CategoryID, &interest.InterestName)
		if err != nil {
			return nil, err
		}
		interests = append(interests, interest)
	}
	return &interests, nil
}

// Get all the interest for the user for Matching and or to show what the user has active
func GetAllUserInterestIDs(userID int) (*[]int, error) {
	query := "SELECT interest_id FROM user_interest WHERE user_interests.user_id = $1"
	rows, err := DB.Query(query, userID)
	if err != nil {
		log.Println("Error getting all user interests_ids")
		return nil, err
	}
	defer rows.Close()
	var interestIDs []int
	for rows.Next() {
		var interestID int
		err = rows.Scan(&interestID)
		if err != nil {
			log.Println("Error scanning row")
			return nil, err
		}
		interestIDs = append(interestIDs, interestID)
	}
	return &interestIDs, nil
}


func AddInterestToUser(interestID int, userID int) error {
	query := "INSERT INTO user_interests (user_id, interest_id) VALUES ($1, $2)"
	_, err := DB.Exec(query, userID, interestID)
	if err != nil {
		log.Println("Error adding interest to user")
		return err
	}
	return nil
}

func RemoveInterestFromUser(interestID int, userID int) error {
	query := "DELETE FROM user_interests WHERE user_id = $1 AND interest_id = $2"
	_, err := DB.Exec(query, userID, interestID)
	if err != nil {
		log.Println("Error removing interest from user")
		return err
	}
	return nil
}
