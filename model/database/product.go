package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"time"
)

type Product struct {
	ID             int       `json:"id" gorm:"primary_key;autoIncrement"`
	Name           string    `json:"name" gorm:"type:varchar(255);not null"`
	Picture        string    `json:"picture" gorm:"type:varchar(255)"`
	Price          float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock          int       `json:"stock" gorm:"type:int;default:0"`
	Description    string    `json:"description" gorm:"type:text"`
	ExpirationTime time.Time `json:"expiration_time" gorm:"type:datetime"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Product) TableName() string {
	return "products"
}

func (product *Product) model() *gorm.DB { return infrastructure.Db.Model(product) }

func (product *Product) Create() error {
	return product.model().Create(product).Error
}

func (product *Product) FindByID(id int) error {
	return product.model().First(product, id).Error
}

func (product *Product) Update(updateData *Product) error {
	return product.model().Updates(updateData).Error
}

func (product *Product) Delete() error {
	return product.model().Delete(product).Error
}

func FindAll() (products []Product, err error) {
	err = infrastructure.Db.Find(&products).Error
	return
}
