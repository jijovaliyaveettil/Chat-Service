package initializers

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

	var err error

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
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Add this to handle foreign key constraints
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Drop existing tables if they exist
	// db.Migrator().DropTable(&models.Friendships{}, &models.User{})

	// AutoMigrate with explicit foreign key configuration
	err = DB.AutoMigrate(&models.User{}, &models.Friendships{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Set up connection pool and other configurations
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	// Optional: Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = DB
	log.Println("Database connection successfully established")
}
