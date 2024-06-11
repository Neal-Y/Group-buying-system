package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"time"
)

type Product struct {
	ID             int       `gorm:"primary_key;autoIncrement"`
	Name           string    `gorm:"type:varchar(255);not null"`
	Picture        string    `gorm:"type:varchar(255)"`
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	Stock          int       `gorm:"type:int;default:0"`
	Description    string    `gorm:"type:text"`
	ExpirationTime time.Time `gorm:"type:timestamp"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
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

func (Product) FindAll() (products []Product, err error) {
	err = infrastructure.Db.Find(&products).Error
	return
}
