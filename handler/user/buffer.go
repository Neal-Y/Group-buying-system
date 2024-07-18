package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *User) Buffer(c *gin.Context) {
	token := c.Query("token")
	displayName := c.Query("display_name")

	c.HTML(http.StatusOK, "buffer.html", gin.H{
		"Authorization": token,
		"display_name":  displayName,
	})
}
