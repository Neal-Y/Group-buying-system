package admin

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/middleware"
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
	adminRoute(h, r)

	return h
}

func loginRoute(h *Admin, r *gin.RouterGroup) {
	r.POST("/admin/login", h.Login)
}

func adminRoute(h *Admin, r *gin.RouterGroup) {
	adminRoute := r.Group("/admin")
	adminRoute.Use(middleware.JWTAuthMiddleware())
	adminRoute.GET("/:id", h.GetAdmin)
	adminRoute.GET("", h.ListAdmins)
	adminRoute.PATCH("/:id", h.UpdateAdmin)
	adminRoute.DELETE("/:id", h.DeleteAdmin)
	adminRoute.POST("/register", h.Register)
}
