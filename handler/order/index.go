package order

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/repository"
	"shopping-cart/service"
	"shopping-cart/util"
)

type Order struct {
	orderService service.OrderService
}

func NewOrderHandler(r *gin.RouterGroup) *Order {
	orderRepo := repository.NewOrderRepository()
	productRepo := repository.NewProductRepository()
	userRepo := repository.NewUserRepository()
	notificationService := service.NewNotificationService()
	notificationCache := util.NewNotificationCache()

	orderService := service.NewOrderService(orderRepo, productRepo, userRepo, notificationService, notificationCache)

	h := &Order{
		orderService: orderService,
	}

	adminRoute(h, r)
	newRoute(h, r)

	return h
}

func newRoute(h *Order, r *gin.RouterGroup) {
	orderRoute := r.Group("/orders")
	orderRoute.Use(middleware.JWTAuthMiddleware(constant.UserType))
	orderRoute.POST("", h.CreateOrder)
	orderRoute.DELETE("/:id", h.DeleteOrder)
	orderRoute.GET("/history", h.ListHistoryOrdersByUser)
}

func adminRoute(h *Order, r *gin.RouterGroup) {
	adminRoute := r.Group("/orders")
	adminRoute.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminRoute.PATCH("/:id", h.UpdateOrder)
	adminRoute.GET("/search", h.SearchOrders)
}
