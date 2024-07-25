package user

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/service"
)

type User struct {
	userService service.UserService
}

func NewAuthorization(r *gin.RouterGroup, userService service.UserService) *User {
	h := &User{
		userService: userService,
	}

	home(h, r)
	lineRoute(h, r)
	manageUser(h, r)
	errorRoute(h, r)
	buffer(h, r)
	emailRoute(h, r)
	resetPasswordRoute(h, r)

	return h
}

func home(h *User, r *gin.RouterGroup) {
	r.GET("/home", h.Home)
}

func buffer(h *User, r *gin.RouterGroup) {
	r.GET("/buffer", h.Buffer)
}

func resetPasswordRoute(h *User, r *gin.RouterGroup) {
	user := r.Group("/user")
	user.POST("/get_email", h.GetUserEmail)
	user.POST("/request_password_reset", h.RequestPasswordReset)
	user.POST("/reset_password", h.ResetPassword)
}

func lineRoute(h *User, r *gin.RouterGroup) {
	r.GET("/line/login", h.LineLogin)
	r.GET("/line/callback", h.LineCallback)
}

func emailRoute(h *User, r *gin.RouterGroup) {
	r.POST("/email/login", h.EmailLogin)
	r.POST("/email/register", h.EmailRegister)
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
