package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "shopping-cart/middlerware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type User struct {
	service service.UserService
}

func NewAuthorization(r *gin.RouterGroup) *User {
	userRepo := repository.NewUserRepository()

	userService := service.NewUserService(userRepo)

	h := &User{
		service: userService,
	}

	newRoute(h, r)

	r.Use(middleware.JWTAuthMiddleware())
	{
		manageUser(h, r)
	}

	return h
}

func newRoute(h *User, r *gin.RouterGroup) {
	r.GET("/line", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}

func manageUser(h *User, r *gin.RouterGroup) {
	r.POST("/admin/users", h.CreateUser)
	r.GET("/admin/users/:id", h.GetUser)
	r.GET("/admin/users", h.GetUsers)
	r.PATCH("/admin/users/:id", h.UpdateUser)
	r.DELETE("/admin/users/:id", h.DeleteUser)
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
