package db

import (
	"fmt"
	"log"
	"match_me_backend/models"
)

func GetAllCategories() (*[]models.Category, error) {
	log.Println("Getting all categories")
	query := "SELECT id, category FROM categories"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Category)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iterations: %w", err)
	}
	return &categories, nil
}
