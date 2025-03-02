package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (change for security)
	},
}

// WebSocket handler
func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New WebSocket connection established")

	// Listen for messages from client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Println("Received:", string(msg))

		// Echo message back to client
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
