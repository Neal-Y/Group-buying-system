package database

import "time"

type VerificationCode struct {
	ID        int       `gorm:"primary_key"`
	Email     string    `gorm:"not null"`
	Code      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Used      bool      `gorm:"default:false"`
}

func (VerificationCode) TableName() string {
	return "verification_codes"
}
