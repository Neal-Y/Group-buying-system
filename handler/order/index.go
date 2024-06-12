package order

import (
	"github.com/gin-gonic/gin"
)

type Order struct{}

func NewOrderHandler(r *gin.RouterGroup) *Order {
	h := &Order{}
	newRoute(h, r)
	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("/orders/:id", h.UpdateOrder)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.GET("/orders", h.ListOrders)
}
