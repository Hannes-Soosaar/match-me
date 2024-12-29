package db

import (
	"log"
	"match_me_backend/models"
)


// type UserInterestResponse struct {
// 	CategoryName string             `json:"category"`
// 	CategoryID   int                `json:"category_id"`
// 	Interests    []models.Interests `json:"interests"`
// }

// func GetInterestResponseBody() (*[]UserInterestResponse, error) {
// 	query := `
// 		SELECT 
// 			c.id AS category_id,
// 			c.category_name,
// 			i.id AS interest_id,
// 			i.category_id,
// 			i.interest_name
// 		FROM 
// 			categories c
// 		LEFT JOIN 
// 			interests i
// 		ON 
// 			c.id = i.category_id
// 		ORDER BY 
// 			c.id, i.id
// 	`

// 	rows, err := DB.Query(query)
// 	if err != nil {
// 		log.Println("Error executing query:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	categoryMap := make(map[int]*UserInterestResponse)

// 	for rows.Next() {
// 		var categoryID int
// 		var categoryName string
// 		var interest models.Interests

// 		err = rows.Scan(&categoryID, &categoryName, &interest.ID, &interest.CategoryID, &interest.InterestName)
// 		if err != nil {
// 			log.Println("Error scanning row:", err)
// 			return nil, err
// 		}

// 		if _, exists := categoryMap[categoryID]; !exists {
// 			categoryMap[categoryID] = &UserInterestResponse{
// 				CategoryName: categoryName,
// 				CategoryID:   categoryID,
// 				Interests:    []models.Interests{},
// 			}
// 		}

// 		if interest.ID != 0 {
// 			categoryMap[categoryID].Interests = append(categoryMap[categoryID].Interests, interest)
// 		}
// 	}


// 	var userInterestResponses []UserInterestResponse
// 	for _, response := range categoryMap {
// 		userInterestResponses = append(userInterestResponses, *response)
// 	}

// 	log.Println("User interest response:", userInterestResponses)
// 	return &userInterestResponses, nil
// }


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


func AddInterestToUser(interestID int, userID string) error {
	query := "INSERT INTO user_interests (user_id, interest_id) VALUES ($1, $2)"
	_, err := DB.Exec(query, userID, interestID)
	if err != nil {
		log.Println("Error adding interest to user")
		return err
	}
	return nil
}

func RemoveInterestFromUser(interestID int, userID string) error {
	query := "DELETE FROM user_interests WHERE user_id = $1 AND interest_id = $2"
	_, err := DB.Exec(query, userID, interestID)
	if err != nil {
		log.Println("Error removing interest from user")
		return err
	}
	return nil
}
