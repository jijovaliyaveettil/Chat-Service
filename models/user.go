package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          string  `gorm:"primaryKey;type:uuid"`
	Name        string  `gorm:"not null"`
	UserName    string  `gorm:"unique;not null"`
	Email       string  `gorm:"unique;not null"`
	Password    string  `gorm:"not null"`
	Friendships []*User `gorm:"many2many:friendships;foreignKey:Id;joinForeignKey:UserID;References:Id;joinReferences:FriendID"`
	CreatedAt   time.Time
}

type Friendships struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    string `gorm:"type:uuid;not null"`
	FriendID  string `gorm:"type:uuid;not null"`
	Status    string `gorm:"size:20;check:status IN ('pending','accepted','rejected')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Chat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	User1     string             `bson:"user1"`    // Sender ID
	User2     string             `bson:"user2"`    // Receiver ID
	Messages  []Message          `bson:"messages"` // Chat messages
	CreatedAt time.Time          `bson:"created_at"`
}

type Message struct {
	SenderID  string    `bson:"sender_id"`
	Content   string    `bson:"content"`
	Timestamp time.Time `bson:"timestamp"`
}
