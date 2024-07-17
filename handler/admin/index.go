package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Admin struct {
	adminService service.AdminService
}

func NewAdminController(r *gin.RouterGroup) *Admin {
	adminRepo := repository.NewAdminRepository()
	verifyRepo := repository.NewVerifyRepository()

	adminService := service.NewAdminService(adminRepo, verifyRepo)

	h := &Admin{
		adminService: adminService,
	}

	loginRoute(h, r)
	adminRoute(h, r)
	resetPasswordRoute(h, r)

	return h
}

func loginRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/login", h.Login)
	r.POST("/register", h.Register) // 只有剛開始的時候才會用到之後會關掉router
}

func adminRoute(h *Admin, r *gin.RouterGroup) {
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.JWTAuthMiddleware(constant.AdminType))
	adminRoute.GET("/:id", h.GetAdmin)
	adminRoute.GET("", h.ListAdmins)
	adminRoute.PATCH("/:id", h.UpdateAdmin)
}

func resetPasswordRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/get_email", h.GetAdminEmail)
	r.POST("/request_password_reset", h.RequestPasswordReset)
	r.POST("/reset_password", h.ResetPassword)
}

func RegisterHomeRoutes(server *gin.Engine) {
	server.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	server.GET("/admin/forgot_password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forgot_password.html", nil)
	})

	server.GET("/admin/reset_password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reset_password.html", nil)
	})

	server.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
