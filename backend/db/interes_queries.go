package db

import "match_me_backend/models"

func GetInterestByUserId(userID int) (*models.Interests, error) {

	return &models.Interests{}, nil
}

// This function will get the Interest based on a query from the FE sending back the category and interest Id

func AddInterestToUser(interestID int, userID int) error {
	return nil
}
//Gets all the interest for FE rendering
func GetAllInterest() (*[]models.Interests, error) {
	return nil, nil
}

// Get all the interest for the user for Matching and or to show what the user has active
func GetAllUserInterestIDs(userID int) (*[]models.Interests, error) {
	return nil, nil
}

// This will return all the Names of the Interest by their ID number
func GetInterestNameByInterestID( interestIDs []int)([]string, error){
	return nil, nil
}