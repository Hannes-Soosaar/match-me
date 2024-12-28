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

// TODO: This function might benefit from being combined to a response in SQL instead of in Go
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
func GetAllUserInterestIDs(userID int) (*[]models.Interests, error) {
	return nil, nil
}

// This will return all the Names of the Interest by their ID number
func GetInterestNameByInterestID(interestIDs []int) ([]string, error) {
	return nil, nil
}

// func InitiateInterestsToUser(userID int) error {
// 	log.Println("Intializing Empty User Interests")
// 	query := "INSERT INTO user_interests (user_id, interest_id) VALUES ($1, $2)"

// 	return nil
// }

//TODO Marko

// This function will get the Interest based on a query from the FE sending back the category and interest Id

func AddInterestToUser(interestID int, userID int) error {
	return nil
}
