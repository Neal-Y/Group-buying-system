package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHomeRoutes(r *gin.Engine) {
	r.GET("/users/login", ShowLogin)
	r.GET("/buffer", ShowBuffer)
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user_login.html", nil)
}

func ShowBuffer(c *gin.Context) {
	c.HTML(http.StatusOK, "buffer.html", nil)
}

//func Error(c *gin.Context) {
//	message := c.Query("message")
//	c.HTML(http.StatusOK, "error.html", gin.H{
//		"errorMessage": message,
//	})
//}
