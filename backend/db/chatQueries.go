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

func SaveNotification(receiverID string, status bool) error {
	query := `
		INSERT INTO unread_messages (receiver, is_unread)
		VALUES ($1, $2)
		ON CONFLICT (receiver)
		DO UPDATE SET is_unread = EXCLUDED.is_unread
	`
	_, err := DB.Exec(query, receiverID, status)
	if err != nil {
		return fmt.Errorf("error saving notification status: %v", err)
	}
	return nil
}

func GetChatHistory(matchID int, offset int, limit int) ([]string, error) {
	query := "SELECT message FROM chat_messages WHERE match_id = $1 ORDER BY sent_at DESC LIMIT $2 OFFSET $3"
	var chatHistory []string

	rows, err := DB.Query(query, matchID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error getting chat history (GetChatHistory): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, fmt.Errorf("error scanning row (GetChatHistory): %v", err)
		}
		chatHistory = append(chatHistory, message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows (GetChatHistory): %v", err)
	}

	return chatHistory, nil
}
