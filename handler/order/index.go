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

	adminRoute(h, r)
	newRoute(h, r)

	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	// add user middleware
	r.POST("/orders", h.CreateOrder)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.GET("/client/history_orders", h.ListHistoryOrdersByUser)
}

func adminRoute(h *Order, r *gin.RouterGroup) {
	adminRoute := r.Group("/orders")
	adminRoute.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminRoute.PATCH("/:id", h.UpdateOrder)
	adminRoute.GET("/search", h.SearchOrders)
}
