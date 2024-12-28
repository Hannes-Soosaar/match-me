package db

import (
	"fmt"
	"log"
)

func SetPicturePath(userID, path string) error {
	userQuery := "UPDATE profiles SET profile_picture = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, path, userID)
	if err != nil {
		log.Printf("Error updating profilke picture for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update the profile picture: %w", err)
	}
	return nil
}
