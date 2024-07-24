package database

import (
	"time"
)

type Product struct {
	ID             int        `json:"id" gorm:"primary_key;autoIncrement"`
	Name           string     `json:"name" gorm:"type:varchar(255);not null;unique"`
	Picture        string     `json:"picture" gorm:"type:varchar(255)"`
	Price          float64    `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock          int        `json:"stock" gorm:"type:int;default:0"`
	Description    string     `json:"description" gorm:"type:text"`
	ExpirationTime time.Time  `json:"expiration_time" gorm:"type:datetime"`
	IsSoldOut      bool       `json:"is_sold_out" gorm:"default:false"`
	SoldOutAt      *time.Time `json:"sold_out_at" gorm:"default:NULL"`
	Supplier       string     `json:"supplier" gorm:"type:varchar(255)"`
}

func (Product) TableName() string {
	return "products"
}
