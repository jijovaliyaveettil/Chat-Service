package handlers

import (
	"chat-service/initializers"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

// WebSocket handler
// func WebSocketHandler(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Println("Failed to upgrade:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	fmt.Println("New WebSocket connection established")

// 	// Listen for messages from client
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Error reading message:", err)
// 			break
// 		}

// 		fmt.Println("Received:", string(msg))

// 		// Echo message back to client
// 		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
// 			fmt.Println("Error writing message:", err)
// 			break
// 		}
// 	}
// }

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // For development only
	},
}

type ChatMessage struct {
	Content string `json:"content"`
}

var (
	clients   = make(map[string]*websocket.Conn)
	clientsMu sync.Mutex
)

func ChatHandler(c *gin.Context) {
	// Get user IDs from query params
	currentUser := c.Query("userID")
	targetUser := c.Query("targetUserID")

	if currentUser == "" || targetUser == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user IDs"})
		return
	}

	// Upgrade to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Register connection
	clientsMu.Lock()
	clients[currentUser] = conn
	clientsMu.Unlock()

	// Set up ping/pong mechanism
	done := make(chan struct{})
	defer close(done)

	// Ping goroutine
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
					log.Println("Ping failed:", err)
					return
				}
			case <-done:
				return
			}
		}
	}()

	// Cleanup on disconnect
	defer func() {
		clientsMu.Lock()
		delete(clients, currentUser)
		clientsMu.Unlock()
	}()

	// Message handling loop
	for {
		var msg ChatMessage

		// Set read deadline
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))

		// Read message properly
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Client %s disconnected unexpectedly", currentUser)
			}
			break
		}

		// Only handle text messages
		if messageType != websocket.TextMessage {
			continue
		}

		log.Printf("Received raw message: %s", message)

		// Parse JSON
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("JSON unmarshal error: %v", err)
			conn.WriteJSON(gin.H{"error": "Invalid message format"})
			continue
		}

		// Validate message content
		if msg.Content == "" || len(msg.Content) > 1000 {
			conn.WriteJSON(gin.H{"error": "Message must be 1-1000 characters"})
			continue
		}

		// Save to MongoDB
		if err := saveMessage(currentUser, targetUser, msg.Content); err != nil {
			log.Println("Database error:", err)
			conn.WriteJSON(gin.H{"error": "Failed to save message"})
			continue
		}

		// Forward to target user
		clientsMu.Lock()
		targetConn, exists := clients[targetUser]
		clientsMu.Unlock()

		if exists {
			if err := targetConn.WriteJSON(map[string]interface{}{
				"sender":  currentUser,
				"content": msg.Content,
				"time":    time.Now().Format(time.RFC3339),
			}); err != nil {
				log.Println("Error forwarding message:", err)
			}
		}
	}
}

func saveMessage(sender, receiver, content string) error {
	collection := initializers.MongoDB.Collection("messages")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{
		"sender":    sender,
		"receiver":  receiver,
		"content":   content,
		"timestamp": time.Now(),
	})

	return err
}
