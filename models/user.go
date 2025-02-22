package models

import "time"

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	UserName  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}
