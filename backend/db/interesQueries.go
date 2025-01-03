package db

import (
	"log"
	"match_me_backend/models"
	"github.com/lib/pq"
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
func GetAllUserInterestIDs(userID string) (*[]int, error) {
	query := "SELECT interest_id FROM user_interests WHERE user_interests.user_id = $1"
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

func GetAllUserInterest(userID string) (*[]models.Interests, error) {
	
	userIdsPtr, err := GetAllUserInterestIDs(userID)
	userIds := *userIdsPtr 
	if err != nil {
		log.Println("Error getting all user interests_ids")
		return nil, err
	}
	
	query := "SELECT id, categoryID, interest FROM interests WHERE id = ANY($1)"
	rows, err := DB.Query(query, pq.Array(userIds))
	if err != nil {
		log.Println("Error getting all user interests")
		return nil, err
	}
	defer rows.Close()
	var interests []models.Interests
	for rows.Next() {
		var interest models.Interests
		err = rows.Scan(&interest.ID, &interest.CategoryID, &interest.InterestName)
		if err != nil {
			log.Println("Error scanning row")
			return nil, err
		}
		interests = append(interests, interest)
	}
	return &interests, nil
}


// TODO recalculate all matches for the user
func AddInterestToUser(interestID int, userID string) error {
	isInterest, err := UserHasInterestByInterestId(interestID, userID)
	if err != nil {
		log.Println("Error getting user interest by interest id")
	}
	if isInterest {
		RemoveInterestFromUser(interestID, userID)
	} else {
		query := "INSERT INTO user_interests (user_id, interest_id) VALUES ($1, $2)"
		_, err := DB.Exec(query, userID, interestID)
		if err != nil {
			return err
		}
	}
	return nil
}

//TODO recalculate all matches for the user
func RemoveInterestFromUser(interestID int, userID string) error {
	query := "DELETE FROM user_interests WHERE user_id = $1 AND interest_id = $2"
	_, err := DB.Exec(query, userID, interestID)
	if err != nil {
		log.Println("Error removing interest from user")
		return err
	}
	return nil
}

func UserHasInterestByInterestId(interestID int, userID string) (bool, error) {
	query := "SELECT * FROM user_interests WHERE user_id = $1 AND interest_id = $2"
	rows, err := DB.Query(query, userID, interestID)
	if err != nil {
		log.Println("Error getting user interest by interest id")
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
