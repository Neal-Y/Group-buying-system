package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
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

func (user *User) model() *gorm.DB { return infrastructure.Db.Model(user) }

func (user *User) Create() error {
	return user.model().Create(user).Error
}

func (user *User) FindById(id int) error {
	return user.model().First(user, id).Error
}

func (user *User) Update(updateData *User) error {
	return user.model().Updates(updateData).Error
}

func (user *User) Delete() error {
	return user.model().Delete(user).Error
}
