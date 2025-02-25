package models

import "time"

type User struct {
	Id          string  `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	UserName    string  `gorm:"unique;not null"`
	Email       string  `gorm:"unique;not null"`
	Password    string  `gorm:"not null"`
	Friendships []*User `gorm:"many2many:friendships;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:FriendID"`
	CreatedAt   time.Time
}

type Friendships struct {
	ID        uint `gorm:"primaryKey"`
	UserID    string
	FriendID  string
	Status    string `gorm:"size:20;check:status IN ('pending','accepted','rejected')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
