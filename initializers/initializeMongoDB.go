package initializers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection instance
var MongoDB *mongo.Database

// ConnectMongo initializes MongoDB connection
func ConnectMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB Connection Error:", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB Ping Error:", err)
	}

	fmt.Println("✅ Connected to MongoDB")

	// Select database
	MongoDB = client.Database("chatapp")
}
