package order

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Order struct {
	orderService service.OrderService
}

func NewOrderHandler(r *gin.RouterGroup) *Order {
	orderRepo := repository.NewOrderRepository()
	productRepo := repository.NewProductRepository()
	userRepo := repository.NewUserRepository()

	orderService := service.NewOrderService(orderRepo, productRepo, userRepo)

	h := &Order{
		orderService: orderService,
	}

	newRoute(h, r)
	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders/:id", h.GetOrder)
	r.PATCH("admin/orders/:id", middleware.JWTAuthMiddleware(constant.AdminType), h.UpdateOrder)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.GET("/orders", h.ListOrders)
	r.GET("/client/history_orders", h.ListHistoryOrdersByUser)
	r.GET("/pick_up_by_display_name", h.GetOrderByUserDisplayName)
}
