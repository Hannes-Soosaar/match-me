package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:3000"
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received message: %s\n", message)

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		//somehow get matchID,senderID, receiverID and save message to database.
	}
}

func ChatDataHandler(w http.ResponseWriter, r *http.Request) {
	senderID := r.URL.Query().Get("senderID")
	matchID := r.URL.Query().Get("matchID")

	if senderID == "" || matchID == "" {
		http.Error(w, "Both senderID and matchID are required", http.StatusBadRequest)
		return
	}

	receiverID, err := db.GetReceiverID(matchID, senderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching receiver ID: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(receiverID)
}

func SaveMessageHandler(w http.ResponseWriter, r *http.Request) {

	var messageData struct {
		MatchID    int    `json:"matchID"`
		SenderID   string `json:"senderID"`
		ReceiverID string `json:"receiverID"`
		Message    string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		log.Printf("ERROR: Failed to save message to database. Error: %v, Arguments: %v", err, messageData)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	err = db.SaveMessage(messageData.Message, messageData.MatchID, messageData.SenderID, messageData.ReceiverID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving message: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Message saved successfully")
}
