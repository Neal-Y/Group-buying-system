package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	middleware "shopping-cart/middlerware"
	"shopping-cart/repository"
	"shopping-cart/service"
)

type Admin struct {
	adminService service.AdminService
}

func NewAdminController(r *gin.RouterGroup) *Admin {
	adminRepo := repository.NewAdminRepository()

	adminService := service.NewAdminService(adminRepo)

	h := &Admin{
		adminService: adminService,
	}

	loginRoute(h, r)

	resetPasswordRoute(h, r)

	r.Use(middleware.JWTAuthMiddleware())
	{
		adminRoute(h, r)
	}

	return h
}

func loginRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/login", h.Login)
}

func adminRoute(h *Admin, r *gin.RouterGroup) {
	r.GET("/admin/:id", h.GetAdmin)
	r.GET("/admins", h.ListAdmins)
	r.PATCH("/admin/:id", h.UpdateAdmin)
	r.DELETE("/admin/:id", h.DeleteAdmin)
	r.POST("/admin/register", h.Register)
}

func resetPasswordRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/get_email", h.GetAdminEmail)
	r.POST("/admin/request_password_reset", h.RequestPasswordReset)
	r.POST("/admin/reset_password", middleware.JWTAuthMiddleware(), h.ResetPassword)
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
