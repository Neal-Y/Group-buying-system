package database

import (
	"time"
)

type Order struct {
	ID           int           `json:"id" gorm:"primary_key;autoIncrement"`
	UserID       int           `json:"user_id" gorm:"not null"`
	TotalPrice   float64       `json:"total_price" gorm:"type:decimal(10,2)"`
	Note         string        `json:"note" gorm:"type:text"`
	Status       string        `json:"status" gorm:"type:enum('pending', 'completed', 'cancelled');default:'pending'"`
	CreatedAt    time.Time     `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	User         *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderDetails []OrderDetail `json:"order_details,omitempty" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "orders"
}
