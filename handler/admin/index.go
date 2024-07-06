package admin

import (
	"github.com/gin-gonic/gin"
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
