package db

import (
	"fmt"
)

func SaveMessage(message string, matchID int, senderID string, receiverID string) error {
	query := "INSERT INTO chat_messages (message, match_id, sender_id, receiver_id) VALUES ($1, $2, $3, $4)"
	_, err := DB.Exec(query, message, matchID, senderID, receiverID)
	if err != nil {
		return fmt.Errorf("error saving message: %v", err)
	}
	return nil
}
