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

type ExternalUser struct {
	ID          int        `json:"id" gorm:"primary_key"`
	DisplayName string     `json:"display_name" gorm:"not null"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone" gorm:"type:varchar(15)"`
	IsMember    bool       `json:"is_member" gorm:"default:false"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"default:null"`
	IsDeleted   bool       `json:"is_deleted" gorm:"default:false"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (ExternalUser) TableName() string {
	return "users"
}
