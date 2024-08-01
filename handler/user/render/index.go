package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserHomeRoutes(r *gin.Engine) {
	r.GET("/home", ShowIndex)
}

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Buffer(c *gin.Context) {
	token := c.Query("token")
	displayName := c.Query("display_name")

	c.HTML(http.StatusOK, "buffer.html", gin.H{
		"Authorization": token,
		"display_name":  displayName,
	})
}

func Error(c *gin.Context) {
	message := c.Query("message")
	c.HTML(http.StatusOK, "error.html", gin.H{
		"errorMessage": message,
	})
}
