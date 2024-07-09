package database

import (
	"time"
)

type Admin struct {
	ID           int       `gorm:"primary_key"`
	Username     string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"unique;not null"`
}

func (Admin) TableName() string {
	return "administrators"
}
