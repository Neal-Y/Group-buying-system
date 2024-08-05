package render

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProductPage(r *gin.Engine) {
	r.GET("/products/:id", ShowIndex)
}

func ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
