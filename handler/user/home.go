package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *User) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
