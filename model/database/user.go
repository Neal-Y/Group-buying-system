package database

type User struct {
	ID          int    `gorm:"primary_key"`
	LineID      string `gorm:"unique;not null"`
	DisplayName string `gorm:"not null"`
	Email       string `gorm:"unique"`
	LineToken   string `gorm:"unique"`
	Phone       string `gorm:"type:varchar(15)"`
	IsMember    bool   `gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}
