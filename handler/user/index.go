package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func RegisterHomeRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	server.GET("/error", func(c *gin.Context) {
		message := c.Query("message")
		c.HTML(http.StatusOK, "error.html", gin.H{
			"errorMessage": message,
		})
	})
}
