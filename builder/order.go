package builder

import (
	"shopping-cart/model/database"
	"time"
)

type OrderBuilder struct {
	order *database.Order
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{order: &database.Order{}}
}

func (b *OrderBuilder) SetUserID(userID int) *OrderBuilder {
	b.order.UserID = userID
	return b
}

func (b *OrderBuilder) SetTotalPrice(totalPrice float64) *OrderBuilder {
	b.order.TotalPrice = totalPrice
	return b
}

func (b *OrderBuilder) SetNote(note string) *OrderBuilder {
	b.order.Note = note
	return b
}

func (b *OrderBuilder) SetStatus(status string) *OrderBuilder {
	b.order.Status = status
	return b
}

func (b *OrderBuilder) SetCreatedAt(createdAt time.Time) *OrderBuilder {
	b.order.CreatedAt = createdAt
	return b
}

func (b *OrderBuilder) SetUpdatedAt(updatedAt time.Time) *OrderBuilder {
	b.order.UpdatedAt = updatedAt
	return b
}

func (b *OrderBuilder) SetOrderDetails(orderDetails []database.OrderDetail) *OrderBuilder {
	b.order.OrderDetails = orderDetails
	return b
}

func (b *OrderBuilder) Build() *database.Order {
	return b.order
}
