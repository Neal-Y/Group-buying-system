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
	User         *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderDetails []OrderDetail `json:"order_details,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        int       `json:"id" gorm:"primary_key;autoIncrement"`
	OrderID   int       `json:"order_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Order     *Order    `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product   *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func (order *Order) model() *gorm.DB { return infrastructure.Db.Model(order) }

func (order *Order) Create() error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	for i := range order.OrderDetails {
		order.OrderDetails[i].CreatedAt = time.Now()
		order.OrderDetails[i].UpdatedAt = time.Now()
	}
	return order.model().Create(order).Error
}

func (order *Order) FindByID(id int) error {
	return order.model().Preload("User").Preload("OrderDetails.Product").First(order, id).Error
}

func (order *Order) Update(updateData *Order) error {
	updateData.UpdatedAt = time.Now()
	for i := range updateData.OrderDetails {
		updateData.OrderDetails[i].UpdatedAt = time.Now()
	}
	return order.model().Model(order).Updates(updateData).Error
}

func (order *Order) Delete() error {
	return order.model().Delete(order).Error
}

func FindAllOrders() ([]Order, error) {
	var orders []Order
	if err := infrastructure.Db.Preload("User").Preload("OrderDetails.Product").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
