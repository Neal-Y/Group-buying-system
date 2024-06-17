package route

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/handler/general"
	"shopping-cart/handler/user"
)

func InitGinServer() (server *gin.Engine, err error) {
	server = GinRouter()
	err = server.Run("127.0.0.1:8080")
	return
}

func GinRouter() (server *gin.Engine) {
	server = gin.New()

	server.LoadHTMLGlob("frontend/*")

	api := server.Group("/api")

	group := server.Group("")

	general.NewGeneral(group)
	user.NewAuthorization(api)

	return server
}
