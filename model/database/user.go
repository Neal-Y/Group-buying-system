package database

import (
	"time"
)

type User struct {
	ID          int       `gorm:"primary_key"`
	LineID      string    `gorm:"unique;not null"`
	DisplayName string    `gorm:"not null"`
	Email       string    `gorm:"unique"`
	LineToken   string    `gorm:"unique"`
	Phone       string    `gorm:"type:varchar(15)"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsMember    bool      `gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}
