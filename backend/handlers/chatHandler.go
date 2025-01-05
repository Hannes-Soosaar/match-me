package handlers

import (
	"fmt"
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
