package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/product"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()

	api := server.Group("/api")
	product.NewProductController(api)

	return server
}
