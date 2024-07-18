package user

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type User struct {
	service service.UserService
}

func NewAuthorization(r *gin.RouterGroup) *User {
	userRepo := repository.NewUserRepository()
	orderRepo := repository.NewOrderRepository()

	userService := service.NewUserService(userRepo, orderRepo)

	h := &User{
		service: userService,
	}

	home(h, r)
	lineRoute(h, r)
	manageUser(h, r)
	errorRoute(h, r)
	buffer(h, r)

	return h
}

func home(h *User, r *gin.RouterGroup) {
	r.GET("/home", middleware.JWTAuthMiddleware(constant.UserType), h.Home)
}

func buffer(h *User, r *gin.RouterGroup) {
	r.GET("/buffer", h.Buffer)
}

func lineRoute(h *User, r *gin.RouterGroup) {
	r.GET("/line/login", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}

func manageUser(h *User, r *gin.RouterGroup) {
	adminGroup := r.Group("/admin/users")
	adminGroup.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminGroup.POST("", h.CreateUser)
	adminGroup.GET("/:id", h.GetUser)
	adminGroup.GET("", h.GetUsers)
	adminGroup.GET("/including_blocked", h.ListBlockedUsers)
	adminGroup.PATCH("/:id", h.UpdateUser)
	adminGroup.DELETE("/:id", h.DeleteUser)
}

func errorRoute(h *User, r *gin.RouterGroup) {
	r.GET("/error", h.Error)
}
