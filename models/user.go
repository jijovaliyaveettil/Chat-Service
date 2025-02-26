package models

import "time"

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
