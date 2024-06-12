package database

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
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
	User         User          `json:"user" gorm:"foreignKey:UserID"`
	OrderDetails []OrderDetail `json:"order_details" gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        int       `json:"id" gorm:"primary_key;autoIncrement"`
	OrderID   int       `json:"order_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Order     Order     `json:"order" gorm:"foreignKey:OrderID"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func (order *Order) model() *gorm.DB { return infrastructure.Db.Model(order) }

func (order *Order) Create() error {
	return order.model().Create(order).Error
}

func (order *Order) FindByID(id int) error {
	return order.model().First(order, id).Error
}

func (order *Order) Update(updateData *Order) error {
	return order.model().Updates(updateData).Error
}

func (order *Order) Delete() error {
	return order.model().Delete(order).Error
}

func FindAllOrders() ([]Order, error) {
	var orders []Order
	err := infrastructure.Db.Preload("OrderDetails").Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}
