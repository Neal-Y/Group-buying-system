package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *User) Error(c *gin.Context) {
	message := c.Query("message")
	c.HTML(http.StatusOK, "error.html", gin.H{
		"errorMessage": message,
	})
}
