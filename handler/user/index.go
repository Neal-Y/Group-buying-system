package user

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Authorization struct {
	service service.UserService
}

func NewAuthorization(r *gin.RouterGroup) *Authorization {
	h := &Authorization{}

	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	h.service = userService

	newRoute(h, r)

	return h
}

func newRoute(h *Authorization, r *gin.RouterGroup) {
	r.GET("/line", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}
