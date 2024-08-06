package repository

import (
	"fmt"
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
	"time"
)

type OrderRepository interface {
	Create(order *database.Order) error
	FindByID(id int) (*database.Order, error)
	Update(order *database.Order) error
	SoftDelete(order *database.Order) error
	BeginTransaction() *gorm.DB
	FindPendingOrdersByUserIDTx(tx *gorm.DB, userID int) ([]database.Order, error)
	FindByUserIDAndProductID(userID, productID int) ([]database.Order, error)
	SearchOrders(keyword string, startDate, endDate time.Time, offset int, limit int) ([]database.OrderWitheTime, int64, error)
	GetRevenueByTimePeriod(startDate, endDate time.Time) (float64, error)
	FindByIDAdmin(id int) (*database.OrderWitheTime, error)
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

func (r *orderRepository) FindByUserIDAndProductID(userID int, productID int) ([]database.Order, error) {
	var orders []database.Order
	err := r.db.Preload("OrderDetails", "product_id = ?", productID).
		Where("user_id = ? AND status != ?", userID, "cancelled").
		Joins("JOIN order_details ON orders.id = order_details.order_id").
		Where("order_details.product_id = ?", productID).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) SearchOrders(keyword string, startDate, endDate time.Time, offset int, limit int) ([]database.OrderWitheTime, int64, error) {
	var orders []database.OrderWitheTime
	var count int64

	query := r.db.Model(&database.OrderWitheTime{}).
		Joins("JOIN users ON users.id = orders.user_id")

	if keyword != "" {
		query = query.Where("orders.note LIKE ? OR orders.status LIKE ? OR users.display_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	if !startDate.IsZero() && !endDate.IsZero() {
		query = query.Where("orders.created_at BETWEEN ? AND ?", startDate, endDate)
	}

	err := query.Count(&count).
		Offset(offset).
		Limit(limit).
		Preload("User").
		Find(&orders).Error

	if err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}

func (r *orderRepository) GetRevenueByTimePeriod(startDate, endDate time.Time) (float64, error) {
	var total float64
	result := r.db.Model(&database.Order{}).
		Select("COALESCE(SUM(total_price), 0)").
		Where("status = ? AND updated_at BETWEEN ? AND ?", "completed", startDate, endDate).
		Scan(&total)

	if result.Error != nil {
		fmt.Printf("Error querying revenue: %v\n", result.Error)
		return 0, result.Error
	}
	return total, nil
}

func (r *orderRepository) FindByIDAdmin(id int) (*database.OrderWitheTime, error) {
	var order database.OrderWitheTime
	err := r.db.Where("orders.id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
