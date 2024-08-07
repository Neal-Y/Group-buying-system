package admin

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/constant"
	"shopping-cart/middleware"
	"shopping-cart/service"
)

type Admin struct {
	adminService service.AdminService
}

func NewAdminController(r *gin.RouterGroup, adminService service.AdminService) *Admin {
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
	adminRoute.GET("", h.GetAdmin)
	adminRoute.PATCH("/:id", h.UpdateAdmin)
}

func resetPasswordRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/get_email", h.GetAdminEmail)
	r.POST("/request_password_reset", h.RequestPasswordReset)
	r.POST("/reset_password", h.ResetPassword)
}
