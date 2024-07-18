package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/admin"
	"shopping-cart/handler/order"
	"shopping-cart/handler/product"
	"shopping-cart/handler/user"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()
	server.Use(gin.Logger())
	server.LoadHTMLGlob("frontend/*")

	admin.RegisterHomeRoutes(server)

	api := server.Group("/api")

	product.NewProductController(api)
	order.NewOrderHandler(api)

	user.NewAuthorization(api)
	admin.NewAdminController(api)

	return server
}
