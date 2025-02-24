package infra

import (
	"chat-service/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Explicitly get environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Validate required environment variables
	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("Missing required database environment variables")
	}

	// Database connection parameters
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	// Open connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{}, &models.Friendships{})

	// Add indexes
	// Ensures friendship pairs are unique regardless of order (A→B == B→A)
	db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_friendships ON friendships (LEAST(user_id, friend_id), GREATEST(user_id, friend_id))`)
	// User ID index
	db.Exec("CREATE INDEX idx_user_id ON friendships (user_id)")
	// Friend ID index
	db.Exec("CREATE INDEX idx_friend_id ON friendships (friend_id)")

	// Set up connection pool and other configurations
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)

	DB = db
	log.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {
	return DB
}
