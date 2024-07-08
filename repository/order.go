package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type OrderRepository interface {
	Create(order *database.Order) error
	FindByID(id int) (*database.Order, error)
	Update(order *database.Order) error
	SoftDelete(order *database.Order) error
	FindAll() ([]database.Order, error)
	BeginTransaction() *gorm.DB
	FindPendingOrdersByUserIDTx(tx *gorm.DB, userID int) ([]database.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		db: infrastructure.Db,
	}
}

func (r *orderRepository) Create(order *database.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id int) (*database.Order, error) {
	var order database.Order
	err := r.db.Preload("User").Preload("OrderDetails.Product").Where("status != ?", "cancelled").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) Update(order *database.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) SoftDelete(order *database.Order) error {
	order.Status = "cancelled"
	return r.db.Save(order).Error
}

func (r *orderRepository) FindAll() ([]database.Order, error) {
	var orders []database.Order
	err := r.db.Preload("User").Preload("OrderDetails.Product").Where("status != ?", "cancelled").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *orderRepository) FindPendingOrdersByUserIDTx(tx *gorm.DB, userID int) ([]database.Order, error) {
	var orders []database.Order
	err := tx.Where("user_id = ? AND status = ?", userID, "pending").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
