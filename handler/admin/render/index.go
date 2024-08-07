package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping-cart/constant"
	"shopping-cart/middleware"
)

func RegisterHomeRoutes(r *gin.Engine) {
	r.GET("/home", middleware.JWTAuthMiddleware(constant.AdminType), ShowIndex)
	r.GET("/admin/login", ShowLogin)
}

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_manage.html", nil)
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
