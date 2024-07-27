package database

type Admin struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	LineID       string `gorm:"unique"`
}

func (Admin) TableName() string {
	return "administrators"
}
