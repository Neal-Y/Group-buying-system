package database

type OrderDetail struct {
	ID        int      `json:"id" gorm:"primary_key;autoIncrement"`
	OrderID   int      `json:"order_id" gorm:"not null"`
	ProductID int      `json:"product_id" gorm:"not null"`
	Quantity  int      `json:"quantity" gorm:"not null"`
	Price     float64  `json:"price" gorm:"type:decimal(10,2)"`
	Order     *Order   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
