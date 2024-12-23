package db

import (
	"fmt"
	"log"
	"time"
)

func SetUsername(userID, username string) error {
	userQuery := "UPDATE profiles SET username = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, username, userID)
	if err != nil {
		log.Printf("Error updating username for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update username: %w", err)
	}

	return nil
}

func SetCity(userID, city, longitude, latitude string) error {
	userQuery := `
	UPDATE users 
	SET 
		user_city = $1, 
		register_location = ST_SetSRID(ST_MakePoint($2, $3), 4326) 
	WHERE uuid = $4
`
	_, err := DB.Exec(userQuery, city, longitude, latitude, userID)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update city: %w", err)
	}

	return nil
}

func SetAbout(userID, about string) error {
	userQuery := "UPDATE profiles SET about_me = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, about, userID)
	if err != nil {
		log.Printf("Error updating about me for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update the about me: %w", err)
	}

	return nil
}

func SetBirthdate(userID string, birthdate time.Time) error {
	userQuery := "UPDATE profiles SET birthdate = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, birthdate, userID)
	if err != nil {
		log.Printf("Error updating birthdate for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update birthdate: %w", err)
	}

	return nil
}
