package handlers

import (
	"chat-service/initializers"
	"chat-service/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 1024,
	WriteBufferSize: 1024 * 1024 * 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (change for security)
	},
	// client sync mutex
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

type ChatMessage struct {
	SenderID string `json:"sender_id"`
	Content  string `json:"content"`
}

// WebSocket Chat Handler
func ChatHandler(c *gin.Context) {
	// Upgrade to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade Error:", err)
		return
	}
	defer conn.Close()

	user1 := c.Query("user1") // Sender
	user2 := c.Query("user2") // Receiver

	// MongoDB collection reference
	collection := initializers.MongoDB.Collection("chats")

	// Infinite loop to keep WebSocket connection open
	for {
		var msg ChatMessage

		// Read incoming message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		fmt.Println("Received message:", msg.Content)

		// Find or create a conversation
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		chat := models.Chat{}
		filter := bson.M{"$or": []bson.M{
			{"user1": user1, "user2": user2},
			{"user1": user2, "user2": user1},
		}}

		err = collection.FindOne(ctx, filter).Decode(&chat)
		if err == mongo.ErrNoDocuments {
			// If no chat exists, create a new one
			chat = models.Chat{
				ID:        primitive.NewObjectID(),
				User1:     user1,
				User2:     user2,
				Messages:  []models.Message{},
				CreatedAt: time.Now(),
			}
		}

		// Append new message
		newMessage := models.Message{
			SenderID:  msg.SenderID,
			Content:   msg.Content,
			Timestamp: time.Now(),
		}
		chat.Messages = append(chat.Messages, newMessage)

		// Update the chat document in MongoDB
		update := bson.M{"$set": bson.M{"messages": chat.Messages}}
		_, err = collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
		if err != nil {
			log.Println("Error updating chat:", err)
			break
		}

		// Send message back to client
		if err := conn.WriteJSON(newMessage); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}
