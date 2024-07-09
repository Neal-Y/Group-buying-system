package database

type Admin struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
}

func (Admin) TableName() string {
	return "administrators"
}
