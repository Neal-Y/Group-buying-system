package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	newRoute(h, r)
	manageUser(h, r)

	return h
}

func newRoute(h *User, r *gin.RouterGroup) {
	r.GET("/line", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}

func manageUser(h *User, r *gin.RouterGroup) {
	adminGroup := r.Group("/admin/users")
	adminGroup.Use(middleware.JWTAuthMiddleware())
	adminGroup.POST("", h.CreateUser)
	adminGroup.GET("/:id", h.GetUser)
	adminGroup.GET("", h.GetUsers)
	adminGroup.GET("/including_blocked", h.ListBlockedUsers)
	adminGroup.PATCH("/:id", h.UpdateUser)
	adminGroup.DELETE("/:id", h.DeleteUser)
}

func RegisterHomeRoutes(server *gin.Engine) {
	server.LoadHTMLGlob("frontend/*")

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
