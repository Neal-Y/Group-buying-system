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

func (b *OrderBuilder) WithUserID(userID int) *OrderBuilder {
	b.order.UserID = userID
	return b
}

func (b *OrderBuilder) WithTotalPrice(totalPrice float64) *OrderBuilder {
	b.order.TotalPrice = totalPrice
	return b
}

func (b *OrderBuilder) WithNote(note string) *OrderBuilder {
	b.order.Note = note
	return b
}

func (b *OrderBuilder) WithStatus(status string) *OrderBuilder {
	b.order.Status = status
	return b
}

func (b *OrderBuilder) WithCreatedAt(createdAt time.Time) *OrderBuilder {
	b.order.CreatedAt = createdAt
	return b
}

func (b *OrderBuilder) WithUpdatedAt(updatedAt time.Time) *OrderBuilder {
	b.order.UpdatedAt = updatedAt
	return b
}

func (b *OrderBuilder) WithOrderDetails(orderDetails []database.OrderDetail) *OrderBuilder {
	b.order.OrderDetails = orderDetails
	return b
}

func (b *OrderBuilder) Build() *database.Order {
	return b.order
}
