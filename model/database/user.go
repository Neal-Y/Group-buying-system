package database

import (
	"time"
)

type User struct {
	ID           int    `gorm:"primary_key"`
	LineID       string `gorm:"not null"`
	DisplayName  string `gorm:"not null"`
	Email        string
	LineToken    string
	Phone        string     `gorm:"type:varchar(15)"`
	IsMember     bool       `gorm:"default:false"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"default:null"`
	IsDeleted    bool       `json:"is_deleted" gorm:"default:false"`
	PasswordHash string     `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "users"
}
