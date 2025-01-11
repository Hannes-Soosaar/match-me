package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*websocket.Conn)
var mu sync.Mutex

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

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Println("No userID provided")
		return
	}

	mu.Lock()
	connections[userID] = conn
	mu.Unlock()
	log.Printf("WebSocket connection established for userID: %s\n", userID)
	log.Printf("Number of connections: %d\n", len(connections))
	log.Printf("Connections: %d\n" , connections)

	defer func() {
		mu.Lock()
		delete(connections, userID)
		mu.Unlock()
		log.Printf("Websocket connection closed for userID: %s\n", userID)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		var msgData struct {
			SenderID   string `json:"senderID"`
			ReceiverID string `json:"receiverID"`
			Message    string `json:"message"`
			Type       string `json:"type"`
		}

		if err := json.Unmarshal(message, &msgData); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		mu.Lock()
		senderConn, senderOnline := connections[msgData.SenderID]
		receiverConn, receiverOnline := connections[msgData.ReceiverID]
		mu.Unlock()

		//returnMessage := string(msgData.Message)

		if msgData.Type == "typing" || msgData.Type == "stopTyping" {
			if receiverOnline {
				err := receiverConn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("Error sending typing status to receiver:", err)
				}
			}
		}

		//if sender is online, send message
		if senderOnline {
			err := senderConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending message to sender:", err)
			}
		}

		//if receiver is online, send message
		if receiverOnline {
			err := receiverConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending message to receiver:", err)
			}
		}
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

func ChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	matchIDStr := r.URL.Query().Get("matchID")
	if matchIDStr == "" {
		http.Error(w, "Missing matchID", http.StatusBadRequest)
		return
	}

	matchID, err := strconv.Atoi(matchIDStr)
	if err != nil {
		http.Error(w, "Invalid matchID", http.StatusBadRequest)
	}

	chatHistory, err := db.GetChatHistory(matchID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting chat history: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatHistory)
}

// I will add this to check if it can be used for online-offline status extraction HS
func GetEstablishedConnections() map[string]*websocket.Conn {
	return connections
}