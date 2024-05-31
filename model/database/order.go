package database

//
//import "time"
//
//// models/order.go
//package models
//
//type Order struct {
//	ID        int     `json:"id" gorm:"primary_key"`
//	UserID    int     `json:"user_id"`
//	ProductID int     `json:"product_id"`
//	Quantity  int     `json:"quantity"`
//	TotalPrice float64 `json:"total_price"`
//	Status    string  `json:"status" gorm:"default:'pending'"`
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}
//
//func CreateOrder(order *Order) error {
//	return db.Create(order).Error
//}
//
//func GetOrderCount(userID string) (int, error) {
//	var count int
//	if err := db.Model(&Order{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
//		return 0, err
//	}
//	return count, nil
//}
